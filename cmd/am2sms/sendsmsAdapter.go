package am2sms

import (
	"os"

	"github.com/StephaneBunel/alertmanager2sms/pkg/appconfig"
	"github.com/StephaneBunel/alertmanager2sms/pkg/domain"
	"github.com/StephaneBunel/alertmanager2sms/pkg/sendsmsadapter"
	"github.com/romana/rlog"
)

// CreateSendsmsAdapter returns a new sendsms library adapter object
func CreateSendsmsAdapter(cnf *appconfig.AppConfig, amEventChan domain.AmEventChan, rrs domain.RecipientRepositoryer, metric domain.Metricer) *sendsmsadapter.Handle {
	cnfKey := cnf.AppName() + ".sms"
	configSms := cnf.Viper.Sub(cnfKey)
	if configSms == nil {
		rlog.Error("Configuration missing:", cnfKey)
		os.Exit(1)
	}
	cnfKey = "sendsms"
	configSendSms := cnf.Viper.Sub(cnfKey)
	if configSendSms == nil {
		rlog.Error("Configuration missing:", cnfKey)
		os.Exit(1)
	}
	ssa := sendsmsadapter.New(configSms, configSendSms, amEventChan, rrs, metric)
	rlog.Debug("Send Sms Interactor:", ssa)
	rlog.Info("sendsms adapter successfully created")
	return ssa
}
