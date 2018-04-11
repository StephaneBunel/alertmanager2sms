package prom

import (
	"github.com/StephaneBunel/alertmanager2sms/pkg/appconfig"
	"github.com/StephaneBunel/alertmanager2sms/pkg/domain"
	"github.com/spf13/viper"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/romana/rlog"
)

func (pm *promMetric) initPromCounters() {
	appName := (&appconfig.AppConfig{}).AppName() + "_"
	pm.promCounterMap = make(map[string]prometheus.Counter)

	for _, item := range []struct {
		name string
		help string
	}{
		{"web_request_total", "Total number of alertmanager events received."},
		{"web_request_decode_error", "Total number of error receiving alertmanager events."},
		{"event_dropped_total", "Number of droped events (event buffer full)."},
		{"sms_sent_total", "Total number of SMS sent."},
	} {
		counter := prometheus.NewCounter(prometheus.CounterOpts{
			Name: appName + item.name,
			Help: item.help,
		})
		prometheus.MustRegister(counter)
		pm.promCounterMap[item.name] = counter
	}
}

func (pm *promMetric) Config(conf *viper.Viper) error {

	pm.once.Do(func() {
		pm.config = conf
		pm.initPromCounters()
		go pm.Serve()
	})

	return nil
}

func (pm *promMetric) IncCounter(name string) {
	if counter, exists := pm.promCounterMap[name]; exists {
		counter.Inc()
	} else {
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
