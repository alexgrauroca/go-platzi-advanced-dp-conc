package chat

func Broadcast() {
	clients := make(map[Client]bool)

	for {
		select {
		case message := <-chatMessages:
			for client := range clients {
				client <- message
			}
		case newClient := <-incomingClients:
			clients[newClient] = true
		case leavingClient := <-leavingClients:
			delete(clients, leavingClient)
			close(leavingClient)
		}
	}
}
