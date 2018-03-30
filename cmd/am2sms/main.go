package main

func main() {
	conf := CreateConfig()
	recipientRepository := CreateRecipientRepository(conf)
	amEventChan := CreateEventChan(conf)
	webAdapter := CreateWebAdapter(conf, amEventChan)
	sendSmsinteractor := CreateSendsmsInteractor(conf, amEventChan, recipientRepository)
	go sendSmsinteractor.Worker()
	go sendSmsinteractor.Worker()
	webAdapter.Serve()
}
