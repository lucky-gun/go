func (m *Manager) Run() {

	for {

		select {

		case client := <-m.Register:

			m.Clients[client] = true

			log.Println("client connected")

		case client := <-m.Unregister:

			if _, ok := m.Clients[client]; ok {

				delete(m.Clients, client)

				close(client.Send)

				log.Println("client disconnected")
			}

		case msg := <-m.Broadcast:

			for client := range m.Clients {

				select {

				case client.Send <- msg:

				default:

					close(client.Send)

					delete(m.Clients, client)
				}
			}
		}
	}
}
