package am2sms

import (
	"os"

	"github.com/StephaneBunel/alertmanager2sms/pkg/appconfig"
	"github.com/StephaneBunel/alertmanager2sms/pkg/domain"
	"github.com/StephaneBunel/alertmanager2sms/pkg/metric/catalog"

	// Import metric handlers
	_ "github.com/StephaneBunel/alertmanager2sms/pkg/metric/handler"

	"github.com/romana/rlog"
)

// CreateMetric returns a new metric according to user preferences
func CreateMetric(cnf *appconfig.AppConfig) domain.Metricer {
	handler := cnf.Viper.GetString(cnf.AppName() + ".stat.handler")
	if handler == "" {
		rlog.Warn("stat handler is nil. Using dummy handler")
		handler = "dummy"
	}
	rlog.Debug("stat.handler =", handler)
	handlerConfig := cnf.Viper.Sub(cnf.AppName() + ".stat.handlerConfig")
	metric := catalog.Default().New(handler)
	if metric == nil {
		os.Exit(1)
	}
	metric.Config(handlerConfig)
	rlog.Debug("metric:", metric)
	rlog.Info("metric (prometheus) successfully created")
	return metric
}
