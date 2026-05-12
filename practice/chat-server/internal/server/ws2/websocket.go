package ws2

import (
	"bufio"
	"fmt"
	"net"
	"net/http"
)

func HandleWebSocket(w http.ResponseWriter, r *http.Request) {

	// 1. websocket key 확인
	key := r.Header.Get("Sec-WebSocket-Key")

	if key == "" {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	// 2. Hijacker 가져오기
	hijacker, ok := w.(http.Hijacker)
	if !ok {
		http.Error(w, "Hijacking not supported", http.StatusInternalServerError)
		return
	}

	// 3. TCP connection 탈취
	conn, rw, err := hijacker.Hijack()
	if err != nil {
		return
	}

	acceptKey := createWebSocketAcceptKey(key)

	// 4. websocket upgrade response
	response := fmt.Sprintf(
		"HTTP/1.1 101 Switching Protocols\r\n"+
			"Upgrade: websocket\r\n"+
			"Connection: Upgrade\r\n"+
			"Sec-WebSocket-Accept: %s\r\n"+
			"\r\n",
		acceptKey,
	)

	_, err = rw.WriteString(response)
	if err != nil {
		conn.Close()
		return
	}

	err = rw.Flush()
	if err != nil {
		conn.Close()
		return
	}

	// 이제부터 raw TCP
	handleConnection(conn)
}

func handleConnection(conn net.Conn) {

	client := &Client{
		Conn: conn,
		Send: make(chan []byte, 256),
	}

	go client.writeLoop()

	client.readLoop()
}
