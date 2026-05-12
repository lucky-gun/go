package model

type ChatPayload struct {
  EventType string `json:"event_type"`

  User    User    `json:"user"`
  Room    Room    `json:"room"`
  Message Message `json:"message"`
}
