package ws

import (
	"github.com/gorilla/websocket"
)

type Client struct {
	Room *Room

	Conn *websocket.Conn

	Send chan []byte
}

func NewClient(room *Room, conn *websocket.Conn) *Client {
	return &Client{
		Room:    room,
		Conn:    conn,
		Send:    make(chan []byte, 256),
	}
}
