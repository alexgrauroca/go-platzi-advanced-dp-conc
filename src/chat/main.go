package chat

import (
	"flag"
	"fmt"
	"log"
	"net"
)

var (
	host = flag.String("h", "localhost", "host")
	port = flag.Int("p", 3090, "port")
)

var (
	incomingClients = make(chan Client)
	leavingClients  = make(chan Client)
	chatMessages    = make(chan string)
)

func Run() {
	flag.Parse()
	connData := fmt.Sprintf("%s:%d", *host, *port)
	listener, err := net.Listen("tcp", connData)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Listening", connData)

	go Broadcast()

	for {
		conn, err := listener.Accept()

		if err != nil {
			log.Print(err)
			continue
		}

		go HandleConnection(conn)
	}
}
