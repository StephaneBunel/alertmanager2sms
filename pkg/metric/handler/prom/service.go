package prom

import (
	"github.com/StephaneBunel/alertmanager2sms/pkg/appconfig"
	"github.com/StephaneBunel/alertmanager2sms/pkg/domain"
	"github.com/spf13/viper"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/romana/rlog"
)

var (
	WebRequestsTotal = prometheus.NewCounter(prometheus.CounterOpts{
		Name: (&appconfig.AppConfig{}).AppName() + "_web_requests_total",
		Help: "Total number of alertmanager events received.",
	})
)

func init() {
	prometheus.MustRegister(WebRequestsTotal)
}

func (pm *promMetric) Config(conf *viper.Viper) error {
	pm.config = conf
	pm.once.Do(func() {
		go pm.Serve()
	})
	return nil
}

func (pm *promMetric) IncCounter(name string) {
	switch name {
	case "web_requests_total":
		WebRequestsTotal.Inc()
		break
	default:
		rlog.Error("metric.IncCounter(", name, "): unknown")
	}
}

func (pm *promMetric) Info() domain.MetricHandlerInfo {
	return domain.MetricHandlerInfo{
		Name:    "prom",
		Authors: "Stéphane Bunel",
		Version: "0.1",
	}
}
