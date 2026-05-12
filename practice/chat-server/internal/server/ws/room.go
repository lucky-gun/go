package ws

import (
  "chat-server/internal/model"
)

type Room struct {
    Room model.Room

    Clients map[*Client]bool

    Register chan *Client
    Unregister chan *Client

    Broadcast chan []byte
}
