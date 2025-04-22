package collector

import (
	"github.com/allsopp/librdkafka-exporter/internal/help"
	"github.com/prometheus/client_golang/prometheus"
)

type Root struct {
	ClientId         string             `json:"client_id"`
	Type             string             `json:"type"`
	Age              float64            `json:"age"`
	MetadataCacheCnt float64            `json:"metadata_cache_cnt"`
	MsgCnt           float64            `json:"msg_cnt"`
	MsgMax           float64            `json:"msg_max"`
	MsgSize          float64            `json:"msg_size"`
	MsgSizeMax       float64            `json:"msg_size_max"`
	ReplyQ           float64            `json:"replyq"`
	RxBytes          float64            `json:"rx_bytes"`
	Rx               float64            `json:"rx"`
	RxMsgBytes       float64            `json:"rxmsg_bytes"`
	RxMsgs           float64            `json:"rxmsgs"`
	SimpleCnt        float64            `json:"simple_cnt"`
	Time             float64            `json:"time"`
	Ts               float64            `json:"ts"`
	TxBytes          float64            `json:"tx_bytes"`
	Tx               float64            `json:"tx"`
	TxMsgBytes       float64            `json:"txmsg_bytes"`
	TxMsgs           float64            `json:"txmsgs"`
	Topics           map[string]*Topic  `json:"topics"`
	Brokers          map[string]*Broker `json:"brokers"`
	Cgrp             *ConsumerGroup     `json:"cgrp"`
}

var getHelp func(string) string = help.Get

func (r *Root) Describe(ch chan<- *prometheus.Desc) {
	prometheus.DescribeByCollect(r, ch)
}

func (r *Root) Collect(ch chan<- prometheus.Metric) {
	var labels = []string{"client_id", "type"}

	ch <- prometheus.MustNewConstMetric(
		prometheus.NewDesc(
			"ts",
			getHelp("ts"),
			labels,
			nil,
		),
		prometheus.CounterValue,
		r.Ts,
		r.ClientId,
		r.Type,
	)
	ch <- prometheus.MustNewConstMetric(
		prometheus.NewDesc("time",
			getHelp("time"),
			labels,
			nil),
		prometheus.CounterValue,
		r.Time,
		r.ClientId,
		r.Type,
	)

	ch <- prometheus.MustNewConstMetric(
		prometheus.NewDesc(
			"age",
			getHelp("age"),
			labels,
			nil,
		),
		prometheus.CounterValue,
		r.Age,
		r.ClientId,
		r.Type,
	)
	ch <- prometheus.MustNewConstMetric(
		prometheus.NewDesc(
			"replyq",
			getHelp("replyq"),
			labels,
			nil,
		),
		prometheus.GaugeValue,
		r.ReplyQ,
		r.ClientId,
		r.Type,
	)
	ch <- prometheus.MustNewConstMetric(
		prometheus.NewDesc(
			"msg_cnt",
			getHelp("msg_cnt"),
			labels,
			nil,
		),
		prometheus.GaugeValue,
		r.MsgCnt,
		r.ClientId,
		r.Type,
	)
	ch <- prometheus.MustNewConstMetric(
		prometheus.NewDesc(
			"msg_size",
			getHelp("msg_size"),
			labels,
			nil,
		),
		prometheus.GaugeValue,
		r.MsgSize,
		r.ClientId,
		r.Type,
	)
	ch <- prometheus.MustNewConstMetric(
		prometheus.NewDesc(
			"msg_max",
			getHelp("msg_max"),
			labels,
			nil,
		),
		prometheus.CounterValue,
		r.MsgMax,
		r.ClientId,
		r.Type,
	)
	ch <- prometheus.MustNewConstMetric(
		prometheus.NewDesc(
			"msg_size_max",
			getHelp("msg_size_max"),
			labels,
			nil,
		),
		prometheus.CounterValue,
		r.MsgSizeMax,
		r.ClientId,
		r.Type,
	)
	ch <- prometheus.MustNewConstMetric(
		prometheus.NewDesc(
			"tx",
			getHelp("tx"),
			labels,
			nil,
		),
		prometheus.CounterValue,
		r.Tx,
		r.ClientId,
		r.Type,
	)
	ch <- prometheus.MustNewConstMetric(
		prometheus.NewDesc(
			"tx_bytes",
			getHelp("tx_bytes"),
			labels,
			nil,
		),
		prometheus.CounterValue,
		r.TxBytes,
		r.ClientId,
		r.Type,
	)
	ch <- prometheus.MustNewConstMetric(
		prometheus.NewDesc(
			"rx",
			getHelp("rx"),
			labels,
			nil,
		),
		prometheus.CounterValue,
		r.Rx,
		r.ClientId,
		r.Type,
	)
	ch <- prometheus.MustNewConstMetric(
		prometheus.NewDesc(
			"rx_bytes",
			getHelp("rx_bytes"),
			labels,
			nil,
		),
		prometheus.CounterValue,
		r.RxBytes,
		r.ClientId,
		r.Type,
	)
	ch <- prometheus.MustNewConstMetric(
		prometheus.NewDesc(
			"txmsgs",
			getHelp("txmsgs"),
			labels,
			nil,
		),
		prometheus.CounterValue,
		r.TxMsgs,
		r.ClientId,
		r.Type,
	)
	ch <- prometheus.MustNewConstMetric(
		prometheus.NewDesc(
			"txmsg_bytes",
			getHelp("txmsg_bytes"),
			labels,
			nil,
		),
		prometheus.CounterValue,
		r.TxMsgBytes,
		r.ClientId,
		r.Type,
	)
	ch <- prometheus.MustNewConstMetric(
		prometheus.NewDesc(
			"rxmsgs",
			getHelp("rxmsgs"),
			labels,
			nil,
		),
		prometheus.CounterValue,
		r.RxMsgs,
		r.ClientId,
		r.Type,
	)
	ch <- prometheus.MustNewConstMetric(
		prometheus.NewDesc(
			"rxmsg_bytes",
			getHelp("rxmsg_bytes"),
			labels,
			nil,
		),
		prometheus.CounterValue,
		r.RxMsgBytes,
		r.ClientId,
		r.Type,
	)
	ch <- prometheus.MustNewConstMetric(
		prometheus.NewDesc(
			"simple_cnt",
			getHelp("simple_cnt"),
			labels,
			nil,
		),
		prometheus.GaugeValue,
		r.SimpleCnt,
		r.ClientId,
		r.Type,
	)
	ch <- prometheus.MustNewConstMetric(
		prometheus.NewDesc(
			"metadata_cache_cnt",
			getHelp("metadata_cache_cnt"),
			labels,
			nil,
		),
		prometheus.GaugeValue,
		r.MetadataCacheCnt,
		r.ClientId,
		r.Type,
	)

	for _, broker := range r.Brokers {
		broker.Collect(ch)
	}

	for _, topic := range r.Topics {
		topic.Collect(ch)
	}

	if r.Cgrp != nil {
		r.Cgrp.Collect(ch)
	}
}
