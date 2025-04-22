package collector

import (
	"os"
	"testing"

	"github.com/prometheus/client_golang/prometheus/testutil"
	"github.com/stretchr/testify/require"
)

func TestRoot(t *testing.T) {
	getHelp = func(string) string { return "" }

	r := &Root{
		ClientId:         "client_id",
		Type:             "type",
		Age:              1,
		MetadataCacheCnt: 2,
		MsgCnt:           3,
		MsgMax:           4,
		MsgSize:          5,
		MsgSizeMax:       6,
		ReplyQ:           7,
		RxBytes:          8,
		Rx:               9,
		RxMsgBytes:       10,
		RxMsgs:           11,
		SimpleCnt:        12,
		Time:             13,
		Ts:               14,
		TxBytes:          15,
		Tx:               16,
		TxMsgBytes:       17,
		TxMsgs:           18,
	}

	fh, err := os.Open("root.example.prom")
	require.Nil(t, err)
	defer fh.Close()

	err = testutil.CollectAndCompare(
		r,
		fh,
		"client_id",
		"type",
		"age",
		"metadata_cache_cnt",
		"msg_cnt",
		"msg_max",
		"msg_size",
		"msg_size_max",
		"replyq",
		"rx_bytes",
		"rx",
		"rxmsg_bytes",
		"rxmsgs",
		"simple_cnt",
		"time",
		"ts",
		"tx_bytes",
		"tx",
		"txmsg_bytes",
		"txmsgs",
	)
	require.Nil(t, err)
}
