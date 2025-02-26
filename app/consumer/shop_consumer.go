package consumer

import (
	"context"
	"fmt"
	"log"

	"github.com/heartlezz7/kafka_practice/config"
	"github.com/segmentio/kafka-go"
)

func ShopConsumer() {

	const (
		topics  = "shop"
		groupID = "shop-group"
	)

	cfg := config.KafkaConnCfg{
		Url:   broker,
		Topic: topics,
	}
	// conn := utils.KafkaConn(cfg)

	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:     []string{cfg.Url},
		Topic:       cfg.Topic,
		GroupID:     groupID,          // Required for automatic acknowledgment
		StartOffset: kafka.LastOffset, // Read only new messages
	})
	defer reader.Close()
	fmt.Println("🚀 Shop Consumer Started (Manual Acknowledgment)...")

	for {
		msg, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Fatalf("❌ Error reading message: %v", err)
		}

		// ✅ Kafka acknowledges this message automatically after reading it
		fmt.Printf("✅ Received: %s\n", string(msg.Value))
	}
}
