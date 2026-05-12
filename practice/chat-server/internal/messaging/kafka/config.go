package kafka

import (
	"os"
	"strings"
)

type Config struct {
	Brokers []string
	Topic   string
	GroupID string
}

func NewConfig() Config {
	brokers := os.Getenv("KAFKA_BROKERS")

	if brokers == "" {
		brokers = "localhost:9092"
	}

	return Config{
		Brokers: strings.Split(brokers, ","),
		Topic:   "chat-topic",
		GroupID: "chat-group",
	}
}
