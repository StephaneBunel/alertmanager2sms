package main

import (
	"os"

	"github.com/StephaneBunel/alertmanager2sms/pkg/appconfig"
	"github.com/StephaneBunel/alertmanager2sms/pkg/domain"
	"github.com/StephaneBunel/alertmanager2sms/pkg/metric/catalog"

	// Import metric handlers
	_ "github.com/StephaneBunel/alertmanager2sms/pkg/metric/handler"

	"github.com/romana/rlog"
)

// CreateWebAdapter returns a new web adapter wich is the REST API
func CreateMetric(cnf *appconfig.AppConfig) domain.Metricer {
	metric := catalog.Default().New("dummy")
	if metric == nil {
		os.Exit(1)
	}
	rlog.Debug("metric:", metric)
	rlog.Info("metric successfully created")
	return metric
}
