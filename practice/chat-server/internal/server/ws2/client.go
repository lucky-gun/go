package websocket_temp

import (
	"log"
	"net"
)

type Client struct {
	Conn net.Conn

	Send chan []byte
}

func (c *Client) readLoop() {

	defer func() {
		log.Println("readLoop closed")
		c.Conn.Close()
	}()

	for {

		payload, err := readFrame(c.Conn)
		if err != nil {
			log.Println("read error:", err)
			return
		}

		log.Println("received:", string(payload))

		// 지금은 echo
		c.Send <- payload
	}
}

func (c *Client) writeLoop() {

	defer func() {
		log.Println("writeLoop closed")
		c.Conn.Close()
	}()

	for {

		msg, ok := <-c.Send

		if !ok {
			return
		}

		err := writeFrame(c.Conn, msg)
		if err != nil {
			log.Println("write error:", err)
			return
		}
	}
}
