package main

import (
	"github.com/StephaneBunel/alertmanager2sms/pkg/appconfig"
	"github.com/StephaneBunel/alertmanager2sms/pkg/domain"
	"github.com/StephaneBunel/alertmanager2sms/pkg/webadapter"
	"github.com/romana/rlog"
)

// WebAdapter
func CreateWebAdapter(cnf *appconfig.AppConfig, amEventChan domain.AmEventChan) *webadapter.WebserviceHandler {
	webAdapter := webadapter.New(cnf.Viper.Sub(cnf.AppName()+".http"), amEventChan)
	rlog.Debug("webadapter:", webAdapter)
	return webAdapter
}
