package chat

import "flag"

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
}
