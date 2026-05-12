package ws

import (
	"log"
	"net/http"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},

       /*CheckOrigin: func(r *http.Request) bool {
       origin := r.Header.Get("Origin")

       return origin == "https://myservice.com"
       }*/
}

func ServeWS(manager *Manager, w http.ResponseWriter, r *http.Request) {

        roomID := r.URL.Query().Get("room")

        if roomID == "" {

        }
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("upgrade error:", err)
		return
	}

	client := NewClient(manager, conn)

	manager.Register <- client

	go client.ReadLoop()
	go client.WriteLoop()
}
