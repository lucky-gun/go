package main

import (
	"log"
        "os"
        "chat-server/internal/messaging"
//	"net/http"
)

func main() {

        log.Println("server start")

        mode := os.Getenv("MQ_TYPE")
        broker, err := messaging.NewMessageBroker(mode)        

        if err != nil {
           log.Fatal(err)
        }

        log.Println("broker connected: %T\n", broker)
        /*
	lobbyManager := lobby.NewLobbyManager()
	roomService := lobby.NewRoomService(lobbyManager)
	eventHandler := lobby.NewEventHandler(roomService)
	wsManager := ws.NewManager(eventHandler)
	http.HandleFunc("/ws", wsManager.HandleWebSocket)
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	log.Println("server started :8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}*/

}
