package consumer

const (
	broker = "localhost:9092"
)

func Consumer() {

	var forever = make(chan bool)

	go ShopConsumer()
	go OrderConsumer()

	<-forever
}
