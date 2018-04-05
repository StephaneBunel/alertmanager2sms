package prom

import (
	"github.com/StephaneBunel/alertmanager2sms/pkg/domain"
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

func (pm *promMetric) Config() {

}

func (pm *promMetric) Info() domain.MetricHandlerInfo {
	return domain.MetricHandlerInfo{
		Name:    "prom",
		Authors: "Stéphane Bunel",
		Version: "0.1",
	}
}
