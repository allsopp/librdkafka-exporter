package collector

import (
	"os"
	"testing"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/testutil"
	"github.com/stretchr/testify/require"
)

type MockWindow struct {
	w Window
}

func (mw *MockWindow) Collect(ch chan<- prometheus.Metric) {
	mw.w.Collect(ch, "foo.bar.baz", []string{"name"}, "qux")
}

func (mw *MockWindow) Describe(ch chan<- *prometheus.Desc) {
	prometheus.DescribeByCollect(mw, ch)
}

func TestWindow(t *testing.T) {
	getHelp = func(string) string {
		return "An example help message"
	}

	mw := &MockWindow{
		Window{
			P50:    10,
			P75:    20,
			P90:    30,
			P95:    40,
			P99:    50,
			P99_99: 60,
			Sum:    1000,
			Count:  10,
		},
	}
	fh, err := os.Open("window.example.prom")
	require.Nil(t, err)
	defer fh.Close()

	err = testutil.CollectAndCompare(mw, fh, "foo_bar_baz")
	require.Nil(t, err)
}
