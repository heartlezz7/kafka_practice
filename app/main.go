package main

import (
	"fmt"
	"os"

	"github.com/heartlezz7/kafka_practice/app/consumer"
)

func main() {
	fmt.Println("Consumer is running")
	consumer.Init()

	os.Exit(1)
	fmt.Println("Consumer is stopped")

}
