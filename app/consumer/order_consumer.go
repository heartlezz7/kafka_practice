package consumer

import (
	"fmt"
	"log"

	"github.com/heartlezz7/kafka_practice/config"
	"github.com/heartlezz7/kafka_practice/pkg/utils"
)

func OrderConsumer() {
	cfg := config.KafkaConnCfg{
		Url:   "localhost:9092",
		Topic: "order",
	}
	conn := utils.KafkaConn(cfg)

	fmt.Println("order consumer started")

	for {
		message, err := conn.ReadMessage(10e3)
		if err != nil {
			break
		}
		fmt.Printf("\n\nmessage: %s \n\n", string(message.Value))

	}

	if err := conn.Close(); err != nil {
		log.Fatal("failed to close connection:", err)
	}
}
