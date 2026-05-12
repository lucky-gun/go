package messaging

type MessageBroker interface {
  Publish(message []byte) error
  Consume(handler func([]byte) error) error
  Close() error
}
