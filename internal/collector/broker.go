package collector

import (
	"github.com/prometheus/client_golang/prometheus"
)

type Broker struct {
	Source         string  `json:"source"`
	Name           string  `json:"name"`
	NodeName       string  `json:"nodename"`
	Connects       float64 `json:"connects"`
	Disconnects    float64 `json:"disconnects"`
	OutBufCnt      float64 `json:"outbuf_cnt"`
	OutBufMsgCnt   float64 `json:"outbuf_msg_cnt"`
	ReqTimeouts    float64 `json:"req_timeouts"`
	RxBytes        float64 `json:"rxbytes"`
	RxCorrIDErrs   float64 `json:"rxcorriderrs"`
	RxErrs         float64 `json:"rxerrs"`
	Rx             float64 `json:"rx"`
	RxIdle         float64 `json:"rxidle"`
	RxPartial      float64 `json:"rxpartial"`
	TxBytes        float64 `json:"txbytes"`
	TxErrs         float64 `json:"txerrs"`
	Tx             float64 `json:"tx"`
	TxIdle         float64 `json:"txidle"`
	TxRetries      float64 `json:"txretries"`
	WaitRespCnt    float64 `json:"waitresp_cnt"`
	WaitRespMsgCnt float64 `json:"waitresp_msg_cnt"`
	Wakeups        float64 `json:"wakeups"`
	ZBufGrow       float64 `json:"zbuf_grow"`
	Req            Request `json:"req"`
	IntLatency     *Window `json:"int_latency"`
	OutBufLatency  *Window `json:"outbuf_latency"`
	RTT            *Window `json:"rtt"`
	Throttle       *Window `json:"throttle"`
}

func (b *Broker) Collect(ch chan<- prometheus.Metric) {
	labels := []string{"name", "nodename", "source"}

	ch <- prometheus.MustNewConstMetric(
		prometheus.NewDesc(
			"brokers_outbuf_cnt",
			getHelp("brokers.outbuf_cnt"),
			labels,
			nil,
		),
		prometheus.GaugeValue,
		b.OutBufCnt,
		b.Name,
		b.NodeName,
		b.Source,
	)

	ch <- prometheus.MustNewConstMetric(
		prometheus.NewDesc(
			"brokers_outbuf_msg_cnt",
			getHelp("brokers.outbuf_msg_cnt"),
			labels,
			nil,
		),
		prometheus.GaugeValue,
		b.OutBufMsgCnt,
		b.Name,
		b.NodeName,
		b.Source,
	)

	ch <- prometheus.MustNewConstMetric(
		prometheus.NewDesc(
			"brokers_waitresp_cnt",
			getHelp("brokers.waitresp_cnt"),
			labels,
			nil,
		),
		prometheus.GaugeValue,
		b.WaitRespCnt,
		b.Name,
		b.NodeName,
		b.Source,
	)

	ch <- prometheus.MustNewConstMetric(
		prometheus.NewDesc(
			"brokers_waitresp_msg_cnt",
			getHelp("brokers.waitresp_msg_cnt"),
			labels,
			nil,
		),
		prometheus.GaugeValue,
		b.WaitRespMsgCnt,
		b.Name,
		b.NodeName,
		b.Source,
	)

	ch <- prometheus.MustNewConstMetric(
		prometheus.NewDesc(
			"brokers_tx",
			getHelp("brokers.tx"),
			labels,
			nil,
		),
		prometheus.CounterValue,
		b.Tx,
		b.Name,
		b.NodeName,
		b.Source,
	)

	ch <- prometheus.MustNewConstMetric(
		prometheus.NewDesc(
			"brokers_txbytes",
			getHelp("brokers.txbytes"),
			labels,
			nil,
		),
		prometheus.CounterValue,
		b.TxBytes,
		b.Name,
		b.NodeName,
		b.Source,
	)

	ch <- prometheus.MustNewConstMetric(
		prometheus.NewDesc(
			"brokers_txerrs",
			getHelp("brokers.txerrs"),
			labels,
			nil,
		),
		prometheus.CounterValue,
		b.TxErrs,
		b.Name,
		b.NodeName,
		b.Source,
	)

	ch <- prometheus.MustNewConstMetric(
		prometheus.NewDesc(
			"brokers_txretries",
			getHelp("brokers.txretries"),
			labels,
			nil,
		),
		prometheus.CounterValue,
		b.TxRetries,
		b.Name,
		b.NodeName,
		b.Source,
	)

	ch <- prometheus.MustNewConstMetric(
		prometheus.NewDesc(
			"brokers_txidle",
			getHelp("brokers.txidle"),
			labels,
			nil,
		),
		prometheus.CounterValue,
		b.TxIdle,
		b.Name,
		b.NodeName,
		b.Source,
	)

	ch <- prometheus.MustNewConstMetric(
		prometheus.NewDesc(
			"brokers_req_timeouts",
			getHelp("brokers.req_timeouts"),
			labels,
			nil,
		),
		prometheus.CounterValue,
		b.ReqTimeouts,
		b.Name,
		b.NodeName,
		b.Source,
	)

	ch <- prometheus.MustNewConstMetric(
		prometheus.NewDesc(
			"brokers_rx",
			getHelp("brokers.rx"),
			labels,
			nil,
		),
		prometheus.CounterValue,
		b.Rx,
		b.Name,
		b.NodeName,
		b.Source,
	)

	ch <- prometheus.MustNewConstMetric(
		prometheus.NewDesc(
			"brokers_rxbytes",
			getHelp("brokers.rxbytes"),
			labels,
			nil,
		),
		prometheus.CounterValue,
		b.RxBytes,
		b.Name,
		b.NodeName,
		b.Source,
	)

	ch <- prometheus.MustNewConstMetric(
		prometheus.NewDesc(
			"brokers_rxerrs",
			getHelp("brokers.rxerrs"),
			labels,
			nil,
		),
		prometheus.CounterValue,
		b.RxErrs,
		b.Name,
		b.NodeName,
		b.Source,
	)

	ch <- prometheus.MustNewConstMetric(
		prometheus.NewDesc(
			"brokers_rxcorriderrs",
			getHelp("brokers.rxcorriderrs"),
			labels,
			nil,
		),
		prometheus.CounterValue,
		b.RxCorrIDErrs,
		b.Name,
		b.NodeName,
		b.Source,
	)

	ch <- prometheus.MustNewConstMetric(
		prometheus.NewDesc(
			"brokers_rxpartial",
			getHelp("brokers.rxpartial"),
			labels,
			nil,
		),
		prometheus.CounterValue,
		b.RxPartial,
		b.Name,
		b.NodeName,
		b.Source,
	)

	ch <- prometheus.MustNewConstMetric(
		prometheus.NewDesc(
			"brokers_rxidle",
			getHelp("brokers.rxidle"),
			labels,
			nil,
		),
		prometheus.CounterValue,
		b.RxIdle,
		b.Name,
		b.NodeName,
		b.Source,
	)

	ch <- prometheus.MustNewConstMetric(
		prometheus.NewDesc(
			"brokers_zbuf_grow",
			getHelp("brokers.zbuf_grow"),
			labels,
			nil,
		),
		prometheus.CounterValue,
		b.ZBufGrow,
		b.Name,
		b.NodeName,
		b.Source,
	)

	ch <- prometheus.MustNewConstMetric(
		prometheus.NewDesc(
			"brokers_wakeups",
			getHelp("brokers.wakeups"),
			labels,
			nil,
		),
		prometheus.CounterValue,
		b.Wakeups,
		b.Name,
		b.NodeName,
		b.Source,
	)

	ch <- prometheus.MustNewConstMetric(
		prometheus.NewDesc(
			"brokers_connects",
			getHelp("brokers.connects"),
			labels,
			nil,
		),
		prometheus.CounterValue,
		float64(b.Connects),
		b.Name,
		b.NodeName,
		b.Source,
	)

	ch <- prometheus.MustNewConstMetric(
		prometheus.NewDesc(
			"brokers_disconnects",
			getHelp("brokers.disconnects"),
			labels,
			nil,
		),
		prometheus.CounterValue,
		float64(b.Disconnects),
		b.Name,
		b.NodeName,
		b.Source,
	)

	b.Req.Collect(
		ch,
		"brokers.req",
		labels,
		b.Name,
		b.NodeName,
		b.Source,
	)

	if b.RTT != nil {
		b.RTT.Collect(
			ch,
			"brokers.rtt",
			labels,
			b.Name,
			b.NodeName,
			b.Source,
		)
	}

	if b.IntLatency != nil {
		b.IntLatency.Collect(
			ch,
			"brokers.int_latency",
			labels,
			b.Name,
			b.NodeName,
			b.Source,
		)
	}

	if b.OutBufLatency != nil {
		b.OutBufLatency.Collect(
			ch,
			"brokers.outbuf_latency",
			labels,
			b.Name,
			b.NodeName,
			b.Source,
		)
	}

	if b.Throttle != nil {
		b.Throttle.Collect(
			ch,
			"brokers.throttle",
			labels,
			b.Name,
			b.NodeName,
			b.Source,
		)
	}
}
