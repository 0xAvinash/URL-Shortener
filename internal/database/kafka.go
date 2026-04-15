package database

import (
	"log"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

var KafkaProducer *kafka.Producer

func NewKafkaProducer() *kafka.Producer {
	producer, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
		"acks":              "all", // durability
		"retries":           3,
		"linger.ms":         5,
	})
	if err != nil {
		log.Fatal("failed to create producer:", err)
	}

	KafkaProducer = producer

	return producer
}
