package collector

import (
	"os"
	"testing"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/testutil"
	"github.com/stretchr/testify/require"
)

type MockRequest struct {
	r Request
}

func (mr *MockRequest) Describe(ch chan<- *prometheus.Desc) {
	prometheus.DescribeByCollect(mr, ch)
}

func (mr *MockRequest) Collect(ch chan<- prometheus.Metric) {
	mr.r.Collect(ch, "foo", []string{"broker"}, "bar")
}

func TestRequest(t *testing.T) {
	getHelp = func(string) string { return "" }

	mr := &MockRequest{
		Request{
			"qux":   1,
			"quux":  2,
			"quuux": 3,
		},
	}

	fh, err := os.Open("request.example.prom")
	require.NoError(t, err)
	defer fh.Close()

	err = testutil.CollectAndCompare(
		mr,
		fh,
		"foo",
	)
	require.NoError(t, err)
}
