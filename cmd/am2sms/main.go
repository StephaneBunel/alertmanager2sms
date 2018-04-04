package main

func main() {
	conf := CreateConfig()
	recipientRepository := UseRecipientRepository(conf)
	amEventChan := CreateEventChan(conf)
	webAdapter := CreateWebAdapter(conf, amEventChan)
	sendSmsAdapter := CreateSendsmsAdapter(conf, amEventChan, recipientRepository)
	go sendSmsAdapter.Worker()
	go sendSmsAdapter.Worker()
	webAdapter.Serve()
}
