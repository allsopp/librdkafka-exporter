package metrics

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"

	"github.com/allsopp/librdkafka-exporter/internal/collector"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/common/expfmt"
)

// Metrics holds internal state and contains no exported fields.
type Metrics struct {
	mutex     *sync.RWMutex
	registry  *prometheus.Registry
	collector prometheus.Collector
}

// The global prefix used for metric names.
const PREFIX = "librdkafka_"

// New returns a pointer to a new instance of Metrics. Returns an error if
// there is a problem registering the metrics with the prometheus.Registry.
func New() (*Metrics, error) {
	m := &Metrics{
		registry:  prometheus.NewRegistry(),
		mutex:     &sync.RWMutex{},
		collector: &collector.Root{},
	}

	r := prometheus.WrapRegistererWithPrefix(PREFIX, m.registry)
	err := r.Register(m)
	if err != nil {
		return nil, fmt.Errorf("error registering metrics: %w", err)
	}

	return m, nil
}

// Describe sends descriptions for collected metrics on the provided channel.
// This method holds a reader lock using a sync.RWMutex while collecting the
// descriptions internally.
func (m Metrics) Describe(ch chan<- *prometheus.Desc) {
	m.collector.Describe(ch)
}

// Collect sends metric values on the provided channel. This method holds a
// reader lock using a sync.RWMutex while collecting the metrics internally.
func (m Metrics) Collect(ch chan<- prometheus.Metric) {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	m.collector.Collect(ch)
}

// Handler returns an http.Handler suitable for serving metrics in the
// Prometheus exposition format.
func (m Metrics) Handler() http.Handler {
	opts := promhttp.HandlerOpts{}
	return promhttp.HandlerFor(m.registry, opts)
}

// ReadFrom reads JSON metrics values from the provided [io.Reader] and updates
// the internal state of metric values. The JSON should be of the format
// returned from the [stats_cb] callback of librdkafka. Returns an io.EOF on
// EOF. This method holds a writer lock using a sync.RWMutex while writing the
// metrics internally.
//
// [stats_cb]: https://github.com/confluentinc/librdkafka/blob/master/STATISTICS.md
func (m Metrics) ReadFrom(r io.Reader) error {
	d := json.NewDecoder(r)
	m.mutex.Lock()
	err := d.Decode(m.collector)
	m.mutex.Unlock()
	if err != nil {
		return fmt.Errorf("error decoding metrics: %w", err)
	}
	return nil
}

// WriteTo writes metric values in Prometheus exposition format to the provided
// [io.Writer]. This method holds a reader lock using a sync.RWMutex while
// reading the metrics internally.
func (m Metrics) WriteTo(w io.Writer) error {
	m.mutex.RLock()
	mfs, err := m.registry.Gather()
	m.mutex.RUnlock()
	if err != nil {
		return fmt.Errorf("error gathering metrics: %w", err)
	}
	for _, mf := range mfs {
		_, err := expfmt.MetricFamilyToText(w, mf)
		if err != nil {
			return fmt.Errorf("error converting metric to text: %w", err)
		}
	}
	return nil
}
