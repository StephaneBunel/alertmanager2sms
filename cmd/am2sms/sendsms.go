package main

import (
	"os"

	"github.com/StephaneBunel/alertmanager2sms/pkg/appconfig"
	"github.com/StephaneBunel/alertmanager2sms/pkg/domain"
	"github.com/StephaneBunel/alertmanager2sms/pkg/sendsmsinteractor"
	"github.com/romana/rlog"
)

func CreateSendsmsInteractor(cnf *appconfig.AppConfig, amEventChan domain.AmEventChan, rrs domain.IRecipientRepository) *sendsmsinteractor.SendSmsInteractor {
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
	ssi := sendsmsinteractor.New(configSms, configSendSms, amEventChan, rrs)
	rlog.Debug("Send Sms Interactor:", ssi)

	return ssi
}
