package consumer

func Consumer() {
	var forever = make(chan bool)

	go ShopConsumer()
	go OrderConsumer()

	<-forever
}
