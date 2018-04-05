package main

import (
	"os"

	"github.com/StephaneBunel/alertmanager2sms/pkg/appconfig"
	"github.com/StephaneBunel/alertmanager2sms/pkg/domain"
	"github.com/StephaneBunel/alertmanager2sms/pkg/recipient/catalog"

	// Import recipient handlers
	_ "github.com/StephaneBunel/alertmanager2sms/pkg/recipient/handler"

	"github.com/romana/rlog"
)

// OpenRecipientRepository returns the selected recipients repository handler
func OpenRecipientRepository(conf *appconfig.AppConfig) domain.RecipientRepositoryer {
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
	csvRecipRepo := catalog.Default().New(engine)
	if csvRecipRepo == nil {
		os.Exit(1)
	}
	err := csvRecipRepo.Config(csvConf)
	if err != nil {
		rlog.Error(err)
		os.Exit(1)
	}
	rlog.Debug("recipientRepository:", csvRecipRepo)
	rlog.Info("recipient repository successfully opened")
	return csvRecipRepo
}
