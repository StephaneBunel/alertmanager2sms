package main

import (
	"github.com/StephaneBunel/alertmanager2sms/cmd/am2sms"

	"github.com/romana/rlog"
)

func main() {
	// bootstrap
	conf := am2sms.CreateConfig()

	metric := am2sms.CreateMetric(conf)

	recipientRepository := am2sms.OpenRecipientRepository(conf)

	amEventChan := am2sms.CreateEventChan(conf, metric)

	webAdapter := am2sms.CreateWebAdapter(conf, amEventChan, metric)

	sendSmsAdapter := am2sms.CreateSendsmsAdapter(conf, amEventChan, recipientRepository, metric)
	workers := conf.Viper.GetInt(conf.AppName() + ".sms.worker")
	if workers < 1 {
		workers = 1
	}
	for ; workers > 0; workers-- {
		go sendSmsAdapter.Worker()
		rlog.Info("sendsms worker launched")
	}

	webAdapter.Serve()
}
