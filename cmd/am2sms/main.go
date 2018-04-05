package main

func main() {
	// bootstrap
	conf := CreateConfig()
	metric := CreateMetric(conf)
	recipientRepository := OpenRecipientRepository(conf)
	amEventChan := CreateEventChan(conf, metric)
	webAdapter := CreateWebAdapter(conf, amEventChan, metric)
	sendSmsAdapter := CreateSendsmsAdapter(conf, amEventChan, recipientRepository, metric)
	go sendSmsAdapter.Worker()
	go sendSmsAdapter.Worker()
	webAdapter.Serve()
}
