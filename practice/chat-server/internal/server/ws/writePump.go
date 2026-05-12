package ws

func (c *Client) WriteLoop() {

	defer c.Conn.Close()

	for {

		msg, ok := <-c.Send
		if !ok {
			return
		}

		err := c.Conn.WriteMessage(
			websocket.TextMessage,
			msg,
		)

		if err != nil {
			log.Println("write error:", err)
			return
		}
	}
}
