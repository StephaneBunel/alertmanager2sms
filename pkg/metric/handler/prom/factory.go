package prom

import (
	"github.com/StephaneBunel/alertmanager2sms/pkg/domain"
	"github.com/StephaneBunel/alertmanager2sms/pkg/metric/catalog"
)

var (
	metrics promMetric
)

// New returns a new prometheus metricer
func New() domain.Metricer {
	p := new(promMetric)
	return p
}

func init() {
	catalog.Default().Register(New, (&promMetric{}).Info())
}
