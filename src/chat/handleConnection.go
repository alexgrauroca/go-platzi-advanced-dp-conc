package chat

import (
	"bufio"
	"fmt"
	"net"
)

func HandleConnection(conn net.Conn) {
	defer conn.Close()

	clientMessages := make(chan string)
	go MessageWrite(conn, clientMessages)

	clientName := conn.RemoteAddr().String()
	clientMessages <- fmt.Sprintf("Welcome to the server, %s\n", clientName)

	chatMessages <- fmt.Sprintf("New client, %s\n", clientName)
	incomingClients <- clientMessages

	inputMessage := bufio.NewScanner(conn)

	for inputMessage.Scan() {
		chatMessages <- fmt.Sprintf("%s: %s\n", clientName, inputMessage.Text())
	}

	leavingClients <- clientMessages
	chatMessages <- fmt.Sprintf("Client left %s", clientName)
}
