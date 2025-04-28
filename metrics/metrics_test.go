package metrics_test

import (
	_ "embed"
	"fmt"
	"os"

	"testing"

	"github.com/allsopp/librdkafka-exporter/metrics"
	"github.com/prometheus/client_golang/prometheus/testutil"
	"github.com/stretchr/testify/require"
)

func TestCollect(t *testing.T) {
	t.Parallel()

	m, err := metrics.New()
	require.NoError(t, err)

	data, err := os.Open("metrics.example.json")
	require.NoError(t, err)
	defer data.Close()

	err = m.ReadFrom(data)
	require.NoError(t, err)

	expected, err := os.Open("metrics.example.prom")
	require.NoError(t, err)
	defer expected.Close()

	err = testutil.CollectAndCompare(
		m,
		expected,
		"age",
		"ts",
		"time",
		"replyq",
		"msg_cnt",
		"msg_size",
		"msg_max",
		"msg_size_max",
		"tx",
		"tx_bytes",
		"rx",
		"rx_bytes",
		"txmsgs",
		"txmsg_bytes",
		"rxmsgs",
		"rxmsg_bytes",
		"simple_cnt",
		"metadata_cache_cnt",
		"brokers_connects",
		"brokers_disconnects",
		"brokers_outbufcnt",
		"brokers_outbuf_msg_cnt",
		"brokers_req_timeouts",
		"brokers_rxbytes",
		"brokers_rxcorriderrs",
		"brokers_rxerrs",
		"brokers_rx",
		"brokers_rxidle",
		"brokers_rxpartial",
		"brokers_txbytes",
		"brokers_txerrs",
		"brokers_tx",
		"brokers_txidle",
		"brokers_txretries",
		"brokers_waitresp_cnt",
		"brokers_waitresp_msg_cnt",
		"brokers_wakeups",
		"brokers_zbuf_grow",
		"cgrp_rebalance_age",
		"cgrp_rebalance_cnt",
		"cgrp_assignment_size",
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
		"req",
		"topics_age",
		"topics_metadata_age",
		"topics_batchsize",
		"topics_batchcnt",
		"brokers_rtt",
		"brokers_int_latency",
		"brokers_outbuf_latency",
		"brokers_throttle",
	)
	if err != nil {
		fmt.Print(err)
	}
	require.NoError(t, err)
}
