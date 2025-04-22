package collector

import (
	"os"
	"testing"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/testutil"
	"github.com/stretchr/testify/require"
)

type MockPartition struct {
	p Partition
}

func (mp *MockPartition) Collect(ch chan<- prometheus.Metric) {
	mp.p.Collect(ch, "example-topic")
}

func (mp *MockPartition) Describe(ch chan<- *prometheus.Desc) {
	prometheus.DescribeByCollect(mp, ch)
}

func TestPartition(t *testing.T) {
	getHelp = func(string) string { return "" }

	mp := &MockPartition{
		Partition{
			Partition:            100,
			AckedMsgId:           1,
			AppOffset:            2,
			CommittedLeaderEpoch: 3,
			CommittedOffset:      4,
			ConsumerLag:          5,
			ConsumerLagStored:    6,
			EofOffset:            7,
			FetchQCnt:            8,
			FetchQSize:           9,
			HiOffset:             10,
			LeaderEpoch:          11,
			LoOffset:             12,
			LsOffset:             13,
			MsgQBytes:            14,
			MsgQCnt:              15,
			Msgs:                 16,
			MsgsInFlight:         17,
			NextAckSeq:           18,
			NextErrSeq:           19,
			NextOffset:           20,
			QueryOffset:          21,
			RxBytes:              22,
			RxMsgs:               23,
			RxVerDrops:           24,
			StoredLeaderEpoch:    25,
			StoredOffset:         26,
			TxBytes:              27,
			TxMsgs:               28,
			XmitMsgQBytes:        29,
			XmitMsgQCnt:          30,
		},
	}

	fh, err := os.Open("partition.example.prom")
	require.Nil(t, err)
	defer fh.Close()

	err = testutil.CollectAndCompare(
		mp,
		fh,
		"partitions_acked_msgid",
		"partitions_app_offset",
		"partitions_committed_leader_epoch",
		"partitions_committed_offset",
		"partitions_consumer_lag",
		"partitions_consumer_lag_stored",
		"partitions_eof_offset",
		"partitions_fetchq_cnt",
		"partitions_fetchq_size",
		"partitions_hi_offset",
		"partitions_leader_epoch",
		"partitions_lo_offset",
		"partitions_ls_offset",
		"partitions_msgq_bytes",
		"partitions_msgq_cnt",
		"partitions_msgs",
		"partitions_msgs_inflight",
		"partitions_next_ack_seq",
		"partitions_next_err_seq",
		"partitions_next_offset",
		"partitions_query_offset",
		"partitions_rxbytes",
		"partitions_rxmsgs",
		"partitions_rx_ver_drops",
		"partitions_stored_leader_epoch",
		"partitions_stored_offset",
		"partitions_txbytes",
		"partitions_txmsgs",
		"partitions_xmit_msgq_bytes",
		"partitions_xmit_msgq_cnt",
	)
	require.Nil(t, err)
}
