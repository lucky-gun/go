package kafka

import (
   "context"
   "log"
   "github.com/segmentio/kafka-go"
)

type Kafka struct {
    writer *kafka.Writer
    reader *kafka.Reader
    
    topic string
    groupID string
}

func New(cfg Config) (*Kafka, error) {

     writer := &kafka.Writer{
       Addr: kafka.TCP(cfg.Brokers...),
       Topic: cfg.Topic,
     }

    reader := kafka.NewReader(kafka.ReaderConfig{
       Brokers: cfg.Brokers,
       Topic: cfg.Topic,
       GroupID: cfg.GroupID,
    })

    conn, err := kafka.Dial("tcp", cfg.Brokers[0])
    if err != nil {
      return nil, err
    }

    conn.Close()
    log.Println("kafka start")
    return &Kafka {
      writer: writer,
      reader: reader,
      topic: cfg.Topic,
      groupID: cfg.GroupID,
    }, nil
}

func (k *Kafka) Publish(message []byte) error {
   err := k.writer.WriteMessages (
      context.Background(),
      kafka.Message{
        Value: message,
      },
   )

   if err != nil {
     log.Println("Publish 실패:", err)
     return err
   }

   return nil
}
func (k *Kafka) Consume(handler func([]byte) error) error {
    go func() {
        for {
            msg, err := k.reader.ReadMessage(context.Background())
            if err != nil {
                log.Println("Kafka consume 실패:", err)
                continue
            }

            err = handler(msg.Value)
            if err != nil {
                log.Println("메시지 처리 실패:", err)
            }
        }
    }()

    return nil
}

func (k *Kafka) Close() error {
    if err := k.writer.Close(); err != nil {
        return err
    }

    if err := k.reader.Close(); err != nil {
        return err
    }

    return nil
}
