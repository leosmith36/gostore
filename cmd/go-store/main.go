package main

import (
	"log"
	"lsmith/go-store/internal/server"
	"net"
)

func main() {
	var (
		err error
		ln net.Listener
	)

	if ln, err = net.Listen("tcp", ":8080"); err != nil {
		log.Fatalf("failed to start listener: %v", err)
	}

  log.Print("started listener")

	for {
		var conn net.Conn
		if conn, err = ln.Accept(); err != nil {
			log.Printf("failed to accept connection: %v", err)
		}

		go server.HandleConnection(conn)
	}
}