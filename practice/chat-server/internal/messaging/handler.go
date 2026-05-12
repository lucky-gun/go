package messaging

import (
  "errors"

  "chat-server/internal/messaging/kafka"
  "chat-server/internal/messaging/rabbitmq"
)

func NewMessageBroker(broker string) (MessageBroker, error) {
  switch broker {

  case "rabbitmq" :
    cfg := rabbitmq.NewConfig()
    return rabbitmq.New(cfg)

  case "kafka":
    cfg := kafka.NewConfig()
    return kafka.New(cfg)

  default:
    return nil, errors.New("이상한 입력값")
  }
}
