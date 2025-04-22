package collector

import (
	"os"
	"testing"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/testutil"
	"github.com/stretchr/testify/require"
)

type MockTopic struct {
	Topic
}

func (mt *MockTopic) Describe(ch chan<- *prometheus.Desc) {
	prometheus.DescribeByCollect(mt, ch)
}

func TestTopic(t *testing.T) {
	getHelp = func(string) string { return "" }

	mt := &MockTopic{
		Topic{
			Topic:       "topic",
			Age:         1,
			MetadataAge: 2,
		},
	}

	fh, err := os.Open("topic.example.prom")
	require.Nil(t, err)
	defer fh.Close()

	err = testutil.CollectAndCompare(
		mt,
		fh,
		"topics_age",
		"topics_metadata_age",
	)
	require.Nil(t, err)
}
