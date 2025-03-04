package main

import (
	"context"
	"log"
	"time"

	"github.com/heartlezz7/kafka_practice/config"
	"github.com/heartlezz7/kafka_practice/models"
	"github.com/heartlezz7/kafka_practice/pkg/utils"
	"github.com/segmentio/kafka-go"
)

func main() {
	// Connection part
	cfg := config.KafkaConnCfg{
		Url:   "localhost:9092",
		Topic: "order",
	}
	conn := utils.KafkaConn(cfg)

	// Check topic if already exists or not
	if !utils.IsTopicAlreadyExists(conn, cfg.Topic) {
		topicConfigs := []kafka.TopicConfig{
			{
				Topic:             cfg.Topic,
				NumPartitions:     1,
				ReplicationFactor: 1,
			},
		}

		err := conn.CreateTopics(topicConfigs...)
		if err != nil {
			panic(err.Error())
		}
	}

	// Mock data
	data := func() []kafka.Message {
		products := []models.Product{
			{
				Id:    "2dc7cf08-e238-4faa-bd5f-f1cfe2e0b565",
				Title: "Coffee",
			},
			{
				Id:    "4c56ec5b-d638-42f2-ae1d-38b6fc6d2122",
				Title: "Tea",
			},
			{
				Id:    "36da5a84-f333-4ecf-a2fe-130c3e8d4ef1",
				Title: "Milk",
			},
		}

		orders := []models.Order{
			{
				Id:     "f7b3f1b1-3b82-4b7b-8b3e-4b4b1f3b7b3b",
				Status: "shipping",
				Item:   products,
			},
		}

		// Convert into kafka.Message{}
		messages := make([]kafka.Message, 0)
		for _, p := range orders {
			messages = append(messages, kafka.Message{
				Value: utils.CompressToJsonBytes(&p),
			})
		}
		return messages
	}()

	// Set timeout

	k := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{cfg.Url},
		Topic:    cfg.Topic,
		Balancer: &kafka.LeastBytes{}})

	k.WriteMessages(context.Background(), kafka.Message{
		Value: []byte("Hello Kafka!"),
		Topic: cfg.Topic,
		Key:   []byte("order-123"),
		Headers: []kafka.Header{
			{Key: "message_type", Value: []byte("OrderPlaced")},
		},
	})
	conn.SetReadDeadline(time.Now().Add(10 * time.Second))
	_, err := conn.WriteMessages(data...)
	if err != nil {
		log.Fatal("failed to write messages:", err)
	}

	// Close connection
	if err := conn.Close(); err != nil {
		log.Fatal("failed to close writer:", err)
	}
}
