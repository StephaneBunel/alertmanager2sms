package prom

import (
	"github.com/StephaneBunel/alertmanager2sms/pkg/domain"
)

var (
	metrics promMetrics
)

func New() domain.Metric {
	p := new(promMetric)
	p.metrics = make(map[string]interface{})

	// p.metric["alerts_received_total"] = prometheus.NewCounter(opts prometheus.CounterOpts)
	return p
}
