package ws

func (c *Client) ReadLoop() {

	defer func() {
		c.Manager.Unregister <- c
		c.Conn.Close()
	}()

	for {

		_, msg, err := c.Conn.ReadMessage()
		if err != nil {
			log.Println("read error:", err)
			return
		}

		log.Println("received:", string(msg))

		// broadcast
		c.Manager.Broadcast <- msg
	}
}
