package server

import (
  "encoding/json"

  "chat-server/internal/messaging"
  "chat-server/internal/model"
)

func PublishPayload(
  broker messaging.MessageBroker,
  payload model.ChatPayload,
) error {
  data, err := json.Marshal(payload)
  if err != nil {
    return err
  }

  return broker.Publish(data)
}

func PayloadHandler(data []byte) error {

    var payload model.ChatPayload

    err := json.Unmarshal(data, &payload)
    if err != nil {
        return err
    }

    return nil
}
