package main

import (
	"github.com/StephaneBunel/alertmanager2sms/pkg/appconfig"
	"github.com/StephaneBunel/alertmanager2sms/pkg/domain"
	"github.com/StephaneBunel/alertmanager2sms/pkg/webadapter"
	"github.com/romana/rlog"
)

// CreateWebAdapter returns a new web adapter wich is the REST API
func CreateWebAdapter(cnf *appconfig.AppConfig, amEventChan domain.AmEventChan, metric domain.Metricer) *webadapter.WebserviceHandle {
	webAdapter := webadapter.New(cnf.Viper.Sub(cnf.AppName()+".http"), amEventChan, metric)
	rlog.Debug("webadapter:", webAdapter)
	rlog.Info("web adapter successfully created")
	return webAdapter
}
