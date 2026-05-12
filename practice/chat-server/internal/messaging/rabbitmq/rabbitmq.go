package rabbitmq

import (
  "log"
  "time"
  "github.com/streadway/amqp"
)

type RabbitMQ struct {
  conn       *amqp.Connection
  channel    *amqp.Channel
  queue      amqp.Queue
  exchange   string
  routingKey string
}

func New(cfg Config) (*RabbitMQ, error) {
        var conn *amqp.Connection
        var err error

        for i := 0; i < 5; i++ {
                conn, err = amqp.Dial(cfg.URL)
                if err == nil {
                        break
                }
                log.Println("RabbitMQ 연결 실패, 재시도...", err)
                time.Sleep(2 * time.Second)
        }
        if err != nil {
                return nil, err
        }

        ch, err := conn.Channel()
        if err != nil {
                return nil, err
        }

        err = ch.ExchangeDeclare(
                cfg.Exchange,
                "direct",
                true,  // durable
                false,
                false,
                false,
                nil,
        )
        if err != nil {
                return nil, err
        }
        // 🔥 Queue 선언 (durable)
        q, err := ch.QueueDeclare(
                cfg.Queue,
                true,  // durable
                false,
                false,
                false,
                nil,
        )
        if err != nil {
                return nil, err
        }

        // 🔥 Binding (exchange → queue 연결)
        err = ch.QueueBind(
                q.Name,
                cfg.RoutingKey,
                cfg.Exchange,
                false,
                nil,
        )
        if err != nil {
                return nil, err
        }

        return &RabbitMQ{
                conn:    conn,
                channel: ch,
                queue: q,
                exchange: cfg.Exchange,
                routingKey: cfg.RoutingKey,
        }, nil
}
func (r *RabbitMQ) Publish(message []byte) error {
        err := r.channel.Publish(
                r.exchange,
                r.routingKey,      // routing key
                false,
                false,
                amqp.Publishing{
                        ContentType:  "text/plain",
                        Body:         message,
                        DeliveryMode: amqp.Persistent, // 🔥 메시지 영속성
                },
        )

        if err != nil {
                log.Println("Publish 실패:", err)
                return err
        }
        return nil
}

func (r *RabbitMQ) Consume(handler func([]byte) error) error {
        msgs, err := r.channel.Consume(
                r.queue.Name,
                "",
                false,
                false,
                false,
                false,
                nil,
        )
        if err != nil {
                return err
        }

        go func() {
                for msg := range msgs {
                        err := handler(msg.Body)

                        if err != nil {
                                log.Println("메시지 처리 실패:", err)
                                msg.Nack(false, true)
                        } else {
                                msg.Ack(false)
                        }
                }
        }()
        return nil
}

func (r *RabbitMQ) Close() error {
        if r.channel != nil {
          if err := r.channel.Close(); err != nil {
            return err
	  }
        }
        if r.conn != nil {
          if err := r.conn.Close(); err != nil {
            return err
          }
        }
      return nil
}
