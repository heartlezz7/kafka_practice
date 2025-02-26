package main

import (
	"fmt"

	"github.com/heartlezz7/kafka_practice/app/consumer"
)

func main() {
	fmt.Println("Consumer is running")
	consumer.Consumer()

	fmt.Println("Consumer is stopped")

}
