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

type Metrics struct {
	mutex     *sync.RWMutex
	registry  *prometheus.Registry
	collector prometheus.Collector
}

const PREFIX = "librdkafka_"

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

func (m Metrics) Describe(ch chan<- *prometheus.Desc) {
	m.collector.Describe(ch)
}

func (m Metrics) Collect(ch chan<- prometheus.Metric) {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	m.collector.Collect(ch)
}

func (m Metrics) Read(rdr io.Reader) error {
	var err error
	data, err := io.ReadAll(rdr)
	if err != nil {
		return err
	}

	m.mutex.Lock()
	err = json.Unmarshal(data, m.collector)
	m.mutex.Unlock()
	if err != nil {
		return err
	}

	return nil
}

func (m Metrics) Write(w io.Writer) (int, error) {
	var total int
	mfs, err := m.registry.Gather()
	if err != nil {
		return total, err
	}
	for _, mf := range mfs {
		var n int
		n, err := expfmt.MetricFamilyToText(w, mf)
		total += n
		if err != nil {
			return total, err
		}
	}
	return total, err
}

func (m Metrics) Handler() http.Handler {
	opts := promhttp.HandlerOpts{}
	return promhttp.HandlerFor(m.registry, opts)
}
