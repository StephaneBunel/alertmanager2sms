package main

import (
	"github.com/romana/rlog"
)

func main() {
	// bootstrap
	conf := CreateConfig()

	metric := CreateMetric(conf)

	recipientRepository := OpenRecipientRepository(conf)

	amEventChan := CreateEventChan(conf, metric)

	webAdapter := CreateWebAdapter(conf, amEventChan, metric)

	sendSmsAdapter := CreateSendsmsAdapter(conf, amEventChan, recipientRepository, metric)
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
