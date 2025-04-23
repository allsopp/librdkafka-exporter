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
		return nil, fmt.Errorf("registration failed: %w", err)
	}

	return m, nil
}

// Describe sends descriptions for collected metrics on the provided channel.
func (m Metrics) Describe(ch chan<- *prometheus.Desc) {
	m.collector.Describe(ch)
}

// Collect sends metric values on the provided channel. Holds a reader lock
// using a sync.RWMutex while collecting the metrics internally.
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

// Read reads JSON metrics values from the provided [io.Reader] and updates the
// internal state of metric values. The JSON should be of the format returned
// from the [stats_cb] callback of librdkafka.  Holds a writer lock using a
// sync.RWMutex while writing the metrics internally.
//
// [stats_cb]: https://github.com/confluentinc/librdkafka/blob/master/STATISTICS.md
func (m Metrics) Read(rdr io.Reader) error {
	var err error
	data, err := io.ReadAll(rdr)
	if err != nil {
		return fmt.Errorf("error reading data: %w", err)
	}

	m.mutex.Lock()
	err = json.Unmarshal(data, m.collector)
	m.mutex.Unlock()
	if err != nil {
		return fmt.Errorf("error unmarshalling JSON: %w", err)
	}

	return nil
}

// Write writes metric values in Prometheus exposition format to the provided
// [io.Writer]. Holds a reader lock using a sync.RWMutex while reading the
// metrics internally.
func (m Metrics) Write(w io.Writer) (int, error) {
	var total int
	m.mutex.RLock()
	mfs, err := m.registry.Gather()
	m.mutex.RUnlock()
	if err != nil {
		return total, fmt.Errorf("error gathering metrics: %w", err)
	}
	for _, mf := range mfs {
		n, err := expfmt.MetricFamilyToText(w, mf)
		total += n
		if err != nil {
			return total, fmt.Errorf("error converting metric to text: %w", err)
		}
	}
	return total, nil
}
