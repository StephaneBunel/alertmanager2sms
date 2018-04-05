package prom

import (
	"github.com/StephaneBunel/alertmanager2sms/pkg/domain"
	"github.com/StephaneBunel/alertmanager2sms/pkg/metric/catalog"
)

var (
	metrics promMetric
)

func New() domain.Metricer {
	p := new(promMetric)
	p.metrics = make(map[string]interface{})
	return p
}

func init() {
	catalog.Default().Register(New, (&promMetric{}).Info())
}
