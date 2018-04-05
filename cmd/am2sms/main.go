package main

func main() {
	// bootstrap
	conf := CreateConfig()
	recipientRepository := OpenRecipientRepository(conf)
	amEventChan := CreateEventChan(conf)
	webAdapter := CreateWebAdapter(conf, amEventChan)
	sendSmsAdapter := CreateSendsmsAdapter(conf, amEventChan, recipientRepository)
	go sendSmsAdapter.Worker()
	go sendSmsAdapter.Worker()
	webAdapter.Serve()
}
