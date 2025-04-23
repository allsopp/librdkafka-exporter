package collector

import (
	"os"
	"testing"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/testutil"
	"github.com/stretchr/testify/require"
)

type MockConsumerGroup struct {
	ConsumerGroup
}

func (mcg *MockConsumerGroup) Describe(ch chan<- *prometheus.Desc) {
	prometheus.DescribeByCollect(mcg, ch)
}

func TestConsumerGroup(t *testing.T) {
	getHelp = func(string) string { return "" }

	mcg := &MockConsumerGroup{
		ConsumerGroup{
			AssignmentSize: 1,
			RebalanceAge:   2,
			RebalanceCnt:   3,
		},
	}

	fh, err := os.Open("consumer_group.example.prom")
	require.NoError(t, err)
	defer fh.Close()

	err = testutil.CollectAndCompare(
		mcg,
		fh,
		"cgrp_rebalance_age",
		"cgrp_rebalance_cnt",
		"cgrp_assignment_size",
	)
	require.NoError(t, err)
}
