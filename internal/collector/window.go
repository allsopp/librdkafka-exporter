package collector

import (
	"strings"

	"github.com/prometheus/client_golang/prometheus"
)

type Window struct {
	P50    float64 `json:"p50"`
	P75    float64 `json:"p75"`
	P90    float64 `json:"p90"`
	P95    float64 `json:"p95"`
	P99    float64 `json:"p99"`
	P99_99 float64 `json:"p99_99"`
	Sum    float64 `json:"sum"`
	Count  uint64  `json:"cnt"`
}

func (w *Window) Collect(
	ch chan<- prometheus.Metric,
	name string,
	variableLabels []string,
	labelValues ...string,
) {
	quantiles := map[float64]float64{
		.5000: w.P50,
		.7500: w.P75,
		.9000: w.P90,
		.9500: w.P95,
		.9900: w.P99,
		.9999: w.P99_99,
	}

	ch <- prometheus.MustNewConstSummary(
		prometheus.NewDesc(
			strings.ReplaceAll(name, ".", "_"),
			getHelp(name),
			variableLabels,
			nil,
		),
		w.Count,
		w.Sum,
		quantiles,
		labelValues...,
	)
}
