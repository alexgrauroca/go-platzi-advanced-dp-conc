package main

import (
	"flag"
	"go-platzi-advanced-dp-conc/src/chat"
	"go-platzi-advanced-dp-conc/src/netcat"
)

var app = flag.String("app", "chat", "Select chat or netcat app to run")

func main() {
	flag.Parse()

	switch *app {
	case "netcat":
		netcat.Run()
	default:
		chat.Run()
	}
}
