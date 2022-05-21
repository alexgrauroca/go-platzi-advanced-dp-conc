package net

import (
	"flag"
	"fmt"
	"net"
	"sync"
)

var site = flag.String("site", "scanme.nmap.org", "Url to scan")

func Run() {
	var wg sync.WaitGroup
	flag.Parse()
	fmt.Println("Site", *site)

	for port := 0; port < 65535; port++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup, port int) {
			defer wg.Done()
			conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", *site, port))

			if err != nil {
				return
			}

			conn.Close()
			fmt.Printf("Port %d opened\n", port)
		}(&wg, port)
	}

	wg.Wait()
}
