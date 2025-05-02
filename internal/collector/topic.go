package collector

import (
	"github.com/prometheus/client_golang/prometheus"
)

type Topic struct {
	Topic       string                `json:"topic"`
	Age         float64               `json:"age"`
	MetadataAge float64               `json:"metadata_age"`
	Partitions  map[string]*Partition `json:"partitions"`
	BatchCnt    *Window               `json:"batchcnt"`
	BatchSize   *Window               `json:"batchsize"`
}

func (t *Topic) Collect(ch chan<- prometheus.Metric) {
	labels := []string{"topic"}

	ch <- prometheus.MustNewConstMetric(
		prometheus.NewDesc(
			"topics_age",
			getHelp("topics.age"),
			labels,
			nil,
		),
		prometheus.GaugeValue,
		t.Age,
		t.Topic,
	)

	ch <- prometheus.MustNewConstMetric(
		prometheus.NewDesc(
			"topics_metadata_age",
			getHelp("topics.metadata_age"),
			labels,
			nil,
		),
		prometheus.GaugeValue,
		t.MetadataAge,
		t.Topic,
	)

	for _, partition := range t.Partitions {
		partition.Collect(ch, t.Topic)
	}

	if t.BatchSize != nil {
		t.BatchSize.Collect(
			ch,
			"topics.batchsize",
			labels,
			t.Topic,
		)
	}

	if t.BatchSize != nil {
		t.BatchCnt.Collect(
			ch,
			"topics.batchcnt",
			labels,
			t.Topic,
		)
	}
}
