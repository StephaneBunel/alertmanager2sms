package prom

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	AlertsReceived = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "alerts_received",
		Help: "Total number of received alert",
	})
)

func init() {

	prometheus.MustRegister(AlertsReceived)
}

func New() IMetric {
	return new(metric)
}

func Inc(metricName string) {

}
