package collector

import (
	"strings"

	"github.com/allsopp/librdkafka-exporter/internal/help"
	"github.com/prometheus/client_golang/prometheus"
)

type Request map[string]float64

func (r Request) Collect(
	ch chan<- prometheus.Metric,
	name string,
	variableLabels []string,
	labelValues ...string,
) {
	for req, count := range r {
		ch <- prometheus.MustNewConstMetric(
			prometheus.NewDesc(
				strings.ReplaceAll(name, ".", "_"),
				help.Get(name),
				variableLabels,
				map[string]string{
					"req": req,
				},
			),
			prometheus.CounterValue,
			count,
			labelValues...,
		)
	}
}
