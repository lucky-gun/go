package rabbitmq

import "os"

type Config struct {
	URL         string
	Exchange    string
	Queue       string
	RoutingKey  string
}

func NewConfig() Config {
	url := os.Getenv("RABBITMQ_URL")

	if url == "" {
		url = "amqp://guest:guest@localhost:5672/"
	}

	return Config{
		URL:         url,
		Exchange:    "chat_exchange",
		Queue:       "chat_queue",
		RoutingKey:  "chat_key",
	}
}
