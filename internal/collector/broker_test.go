package collector

import (
	"os"
	"testing"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/testutil"
	"github.com/stretchr/testify/require"
)

type MockBroker struct {
	Broker
}

func (mcg *MockBroker) Describe(ch chan<- *prometheus.Desc) {
	prometheus.DescribeByCollect(mcg, ch)
}

func TestBroker(t *testing.T) {
	getHelp = func(string) string { return "" }

	mcg := &MockBroker{
		Broker{
			Source:         "source",
			Name:           "name",
			NodeName:       "node",
			Connects:       1,
			Disconnects:    2,
			OutBufCnt:      3,
			OutBufMsgCnt:   4,
			ReqTimeouts:    5,
			RxBytes:        6,
			RxCorrIDErrs:   7,
			RxErrs:         8,
			Rx:             9,
			RxIdle:         10,
			RxPartial:      11,
			TxBytes:        12,
			TxErrs:         13,
			Tx:             14,
			TxIdle:         15,
			TxRetries:      16,
			WaitRespCnt:    17,
			WaitRespMsgCnt: 18,
			Wakeups:        19,
			ZBufGrow:       20,
		},
	}

	fh, err := os.Open("broker.example.prom")
	require.NoError(t, err)
	defer fh.Close()

	err = testutil.CollectAndCompare(
		mcg,
		fh,
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
	)
	require.NoError(t, err)
}
