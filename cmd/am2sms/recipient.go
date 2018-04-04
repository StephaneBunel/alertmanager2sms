package main

import (
	"os"

	"github.com/StephaneBunel/alertmanager2sms/pkg/appconfig"
	"github.com/StephaneBunel/alertmanager2sms/pkg/domain"
	"github.com/StephaneBunel/alertmanager2sms/pkg/recipient"

	// Import recipient plugins
	_ "github.com/StephaneBunel/alertmanager2sms/pkg/recipient/fromcsv"

	"github.com/romana/rlog"
)

func UseRecipientRepository(conf *appconfig.AppConfig) domain.IRecipientRepositoryer {
	keyConf := conf.AppName() + ".recipients.enabled"
	enabledRecipientRepositories := conf.Viper.GetStringSlice(keyConf)
	if len(enabledRecipientRepositories) == 0 {
		rlog.Error("No recipient repository was enabled.", keyConf)
		os.Exit(1)
	}
	rlog.Debug("recipient repository enabled:", enabledRecipientRepositories)
	rr := enabledRecipientRepositories[0]
	subTree := conf.AppName() + ".recipients.config." + rr
	csvConf := conf.Viper.Sub(subTree)
	if csvConf == nil {
		rlog.Error("No configuration for recipient repository", rr, subTree)
		os.Exit(1)
	}
	engine := csvConf.GetString("engine")
	csvRecipRepo := recipient.RepositoryCatalog().New(engine)
	if csvRecipRepo == nil {
		rlog.Error("Recipients repository CSV not found!")
		os.Exit(1)
	}

	err := csvRecipRepo.Config(csvConf)
	if err != nil {
		rlog.Error(err)
		os.Exit(1)
	}

	rlog.Debug("recipientRepository:", csvRecipRepo)
	return csvRecipRepo
}
