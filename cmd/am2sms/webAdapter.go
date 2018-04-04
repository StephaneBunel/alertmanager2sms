package main

import (
	"github.com/StephaneBunel/alertmanager2sms/pkg/appconfig"
	"github.com/StephaneBunel/alertmanager2sms/pkg/domain"
	"github.com/StephaneBunel/alertmanager2sms/pkg/webadapter"
	"github.com/romana/rlog"
)

// CreateWebAdapter returns a new web adapter wich is the REST API
func CreateWebAdapter(cnf *appconfig.AppConfig, amEventChan domain.AmEventChan) *webadapter.WebserviceHandler {
	webAdapter := webadapter.New(cnf.Viper.Sub(cnf.AppName()+".http"), amEventChan)
	rlog.Debug("webadapter:", webAdapter)
	return webAdapter
}
