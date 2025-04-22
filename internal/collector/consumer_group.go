package collector

import (
	"github.com/prometheus/client_golang/prometheus"
)

type ConsumerGroup struct {
	AssignmentSize float64 `json:"assignment_size"`
	RebalanceAge   float64 `json:"rebalance_age"`
	RebalanceCnt   float64 `json:"rebalance_cnt"`
}

func (c *ConsumerGroup) Collect(ch chan<- prometheus.Metric) {
	ch <- prometheus.MustNewConstMetric(
		prometheus.NewDesc(
			"cgrp_rebalance_age",
			getHelp("cgrp.rebalance_age"),
			nil,
			nil,
		),
		prometheus.GaugeValue,
		c.RebalanceAge,
	)

	ch <- prometheus.MustNewConstMetric(
		prometheus.NewDesc(
			"cgrp_rebalance_cnt",
			getHelp("cgrp.rebalance_cnt"),
			nil,
			nil,
		),
		prometheus.CounterValue,
		c.RebalanceCnt,
	)

	ch <- prometheus.MustNewConstMetric(
		prometheus.NewDesc(
			"cgrp_assignment_size",
			getHelp("cgrp.assignment_size"),
			nil,
			nil,
		),
		prometheus.GaugeValue,
		c.AssignmentSize,
	)
}
