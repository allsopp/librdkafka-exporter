package collector

import (
	"fmt"

	"github.com/prometheus/client_golang/prometheus"
)

type Partition struct {
	Partition            float64 `json:"partition"`
	AckedMsgID           float64 `json:"acked_msgid"`
	AppOffset            float64 `json:"app_offset"`
	CommittedLeaderEpoch float64 `json:"committed_leader_epoch"`
	CommittedOffset      float64 `json:"committed_offset"`
	ConsumerLag          float64 `json:"consumer_lag"`
	ConsumerLagStored    float64 `json:"consumer_lag_stored"`
	EOFOffset            float64 `json:"eof_offset"`
	FetchQCnt            float64 `json:"fetchq_cnt"`
	FetchQSize           float64 `json:"fetchq_size"`
	HiOffset             float64 `json:"hi_offset"`
	LeaderEpoch          float64 `json:"leader_epoch"`
	LoOffset             float64 `json:"lo_offset"`
	LsOffset             float64 `json:"ls_offset"`
	MsgQBytes            float64 `json:"msgq_bytes"`
	MsgQCnt              float64 `json:"msgq_cnt"`
	Msgs                 float64 `json:"msgs"`
	MsgsInFlight         float64 `json:"msgs_inflight"`
	NextAckSeq           float64 `json:"next_ack_seq"`
	NextErrSeq           float64 `json:"next_err_seq"`
	NextOffset           float64 `json:"next_offset"`
	QueryOffset          float64 `json:"query_offset"`
	RxBytes              float64 `json:"rxbytes"`
	RxMsgs               float64 `json:"rxmsgs"`
	RxVerDrops           float64 `json:"rx_ver_drops"`
	StoredLeaderEpoch    float64 `json:"stored_leader_epoch"`
	StoredOffset         float64 `json:"stored_offset"`
	TxBytes              float64 `json:"txbytes"`
	TxMsgs               float64 `json:"txmsgs"`
	XmitMsgQBytes        float64 `json:"xmit_msgq_bytes"`
	XmitMsgQCnt          float64 `json:"xmit_msgq_cnt"`
}

func (p *Partition) Collect(ch chan<- prometheus.Metric, topic string) {
	labels := []string{"topic", "partition"}

	ch <- prometheus.MustNewConstMetric(
		prometheus.NewDesc(
			"partitions_msgq_cnt",
			getHelp("partitions.msgq_cnt"),
			labels,
			nil,
		),
		prometheus.GaugeValue,
		p.MsgQCnt,
		topic,
		fmt.Sprint(p.Partition),
	)

	ch <- prometheus.MustNewConstMetric(
		prometheus.NewDesc(
			"partitions_msgq_bytes",
			getHelp("partitions.msgq_bytes"),
			labels,
			nil,
		),
		prometheus.GaugeValue,
		p.MsgQBytes,
		topic,
		fmt.Sprint(p.Partition),
	)

	ch <- prometheus.MustNewConstMetric(
		prometheus.NewDesc(
			"partitions_xmit_msgq_cnt",
			getHelp("partitions.xmit_msgq_cnt"),
			labels,
			nil,
		),
		prometheus.GaugeValue,
		p.XmitMsgQCnt,
		topic,
		fmt.Sprint(p.Partition),
	)

	ch <- prometheus.MustNewConstMetric(
		prometheus.NewDesc(
			"partitions_xmit_msgq_bytes",
			getHelp("partitions.xmit_msgq_bytes"),
			labels,
			nil,
		),
		prometheus.GaugeValue,
		p.XmitMsgQBytes,
		topic,
		fmt.Sprint(p.Partition),
	)

	ch <- prometheus.MustNewConstMetric(
		prometheus.NewDesc(
			"partitions_fetchq_cnt",
			getHelp("partitions.fetchq_cnt"),
			labels,
			nil,
		),
		prometheus.GaugeValue,
		p.FetchQCnt,
		topic,
		fmt.Sprint(p.Partition),
	)

	ch <- prometheus.MustNewConstMetric(
		prometheus.NewDesc(
			"partitions_fetchq_size",
			getHelp("partitions.fetchq_size"),
			labels,
			nil,
		),
		prometheus.GaugeValue,
		p.FetchQSize,
		topic,
		fmt.Sprint(p.Partition),
	)

	ch <- prometheus.MustNewConstMetric(
		prometheus.NewDesc(
			"partitions_query_offset",
			getHelp("partitions.query_offset"),
			labels,
			nil,
		),
		prometheus.GaugeValue,
		p.QueryOffset,
		topic,
		fmt.Sprint(p.Partition),
	)

	ch <- prometheus.MustNewConstMetric(
		prometheus.NewDesc(
			"partitions_next_offset",
			getHelp("partitions.next_offset"),
			labels,
			nil,
		),
		prometheus.GaugeValue,
		p.NextOffset,
		topic,
		fmt.Sprint(p.Partition),
	)

	ch <- prometheus.MustNewConstMetric(
		prometheus.NewDesc(
			"partitions_app_offset",
			getHelp("partitions.app_offset"),
			labels,
			nil,
		),
		prometheus.GaugeValue,
		p.AppOffset,
		topic,
		fmt.Sprint(p.Partition),
	)

	ch <- prometheus.MustNewConstMetric(
		prometheus.NewDesc(
			"partitions_stored_offset",
			getHelp("partitions.stored_offset"),
			labels,
			nil,
		),
		prometheus.GaugeValue,
		p.StoredOffset,
		topic,
		fmt.Sprint(p.Partition),
	)

	ch <- prometheus.MustNewConstMetric(
		prometheus.NewDesc(
			"partitions_stored_leader_epoch",
			getHelp("partitions.stored_leader_epoch"),
			labels,
			nil,
		),
		prometheus.CounterValue,
		p.StoredLeaderEpoch,
		topic,
		fmt.Sprint(p.Partition),
	)

	ch <- prometheus.MustNewConstMetric(
		prometheus.NewDesc(
			"partitions_committed_offset",
			getHelp("partitions.committed_offset"),
			labels,
			nil,
		),
		prometheus.GaugeValue,
		p.CommittedOffset,
		topic,
		fmt.Sprint(p.Partition),
	)

	ch <- prometheus.MustNewConstMetric(
		prometheus.NewDesc(
			"partitions_committed_leader_epoch",
			getHelp("partitions.committed_leader_epoch"),
			labels,
			nil,
		),
		prometheus.CounterValue,
		p.CommittedLeaderEpoch,
		topic,
		fmt.Sprint(p.Partition),
	)

	ch <- prometheus.MustNewConstMetric(
		prometheus.NewDesc(
			"partitions_eof_offset",
			getHelp("partitions.eof_offset"),
			labels,
			nil,
		),
		prometheus.GaugeValue,
		p.EOFOffset,
		topic,
		fmt.Sprint(p.Partition),
	)

	ch <- prometheus.MustNewConstMetric(
		prometheus.NewDesc(
			"partitions_lo_offset",
			getHelp("partitions.lo_offset"),
			labels,
			nil,
		),
		prometheus.GaugeValue,
		p.LoOffset,
		topic,
		fmt.Sprint(p.Partition),
	)

	ch <- prometheus.MustNewConstMetric(
		prometheus.NewDesc(
			"partitions_hi_offset",
			getHelp("partitions.hi_offset"),
			labels,
			nil,
		),
		prometheus.GaugeValue,
		p.HiOffset,
		topic,
		fmt.Sprint(p.Partition),
	)

	ch <- prometheus.MustNewConstMetric(
		prometheus.NewDesc(
			"partitions_ls_offset",
			getHelp("partitions.ls_offset"),
			labels,
			nil,
		),
		prometheus.GaugeValue,
		p.LsOffset,
		topic,
		fmt.Sprint(p.Partition),
	)

	ch <- prometheus.MustNewConstMetric(
		prometheus.NewDesc(
			"partitions_consumer_lag",
			getHelp("partitions.consumer_lag"),
			labels,
			nil,
		),
		prometheus.GaugeValue,
		p.ConsumerLag,
		topic,
		fmt.Sprint(p.Partition),
	)

	ch <- prometheus.MustNewConstMetric(
		prometheus.NewDesc(
			"partitions_consumer_lag_stored",
			getHelp("partitions.consumer_lag_stored"),
			labels,
			nil,
		),
		prometheus.GaugeValue,
		p.ConsumerLagStored,
		topic,
		fmt.Sprint(p.Partition),
	)

	ch <- prometheus.MustNewConstMetric(
		prometheus.NewDesc(
			"partitions_leader_epoch",
			getHelp("partitions.leader_epoch"),
			labels,
			nil,
		),
		prometheus.CounterValue,
		p.LeaderEpoch,
		topic,
		fmt.Sprint(p.Partition),
	)

	ch <- prometheus.MustNewConstMetric(
		prometheus.NewDesc(
			"partitions_txmsgs",
			getHelp("partitions.txmsgs"),
			labels,
			nil,
		),
		prometheus.CounterValue,
		p.TxMsgs,
		topic,
		fmt.Sprint(p.Partition),
	)

	ch <- prometheus.MustNewConstMetric(
		prometheus.NewDesc(
			"partitions_txbytes",
			getHelp("partitions.txbytes"),
			labels,
			nil,
		),
		prometheus.CounterValue,
		p.TxBytes,
		topic,
		fmt.Sprint(p.Partition),
	)

	ch <- prometheus.MustNewConstMetric(
		prometheus.NewDesc(
			"partitions_rxmsgs",
			getHelp("partitions.rxmsgs"),
			labels,
			nil,
		),
		prometheus.CounterValue,
		p.RxMsgs,
		topic,
		fmt.Sprint(p.Partition),
	)

	ch <- prometheus.MustNewConstMetric(
		prometheus.NewDesc(
			"partitions_rxbytes",
			getHelp("partitions.rxbytes"),
			labels,
			nil,
		),
		prometheus.CounterValue,
		p.RxBytes,
		topic,
		fmt.Sprint(p.Partition),
	)

	ch <- prometheus.MustNewConstMetric(
		prometheus.NewDesc(
			"partitions_msgs",
			getHelp("partitions.msgs"),
			labels,
			nil,
		),
		prometheus.CounterValue,
		p.Msgs,
		topic,
		fmt.Sprint(p.Partition),
	)

	ch <- prometheus.MustNewConstMetric(
		prometheus.NewDesc(
			"partitions_rx_ver_drops",
			getHelp("partitions.rx_ver_drops"),
			labels,
			nil,
		),
		prometheus.CounterValue,
		p.RxVerDrops,
		topic,
		fmt.Sprint(p.Partition),
	)

	ch <- prometheus.MustNewConstMetric(
		prometheus.NewDesc(
			"partitions_msgs_inflight",
			getHelp("partitions.msgs_inflight"),
			labels,
			nil,
		),
		prometheus.GaugeValue,
		p.MsgsInFlight,
		topic,
		fmt.Sprint(p.Partition),
	)

	ch <- prometheus.MustNewConstMetric(
		prometheus.NewDesc(
			"partitions_next_ack_seq",
			getHelp("partitions.next_ack_seq"),
			labels,
			nil,
		),
		prometheus.GaugeValue,
		p.NextAckSeq,
		topic,
		fmt.Sprint(p.Partition),
	)

	ch <- prometheus.MustNewConstMetric(
		prometheus.NewDesc(
			"partitions_next_err_seq",
			getHelp("partitions.next_err_seq"),
			labels,
			nil,
		),
		prometheus.GaugeValue,
		p.NextErrSeq,
		topic,
		fmt.Sprint(p.Partition),
	)

	ch <- prometheus.MustNewConstMetric(
		prometheus.NewDesc(
			"partitions_acked_msgid",
			getHelp("partitions.acked_msgid"),
			labels,
			nil,
		),
		prometheus.CounterValue,
		float64(p.AckedMsgID),
		topic,
		fmt.Sprint(p.Partition),
	)
}
