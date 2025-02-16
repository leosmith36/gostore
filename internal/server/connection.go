package server

import (
	"bufio"
	"log"
	"lsmith/go-store/internal/constants"
	"net"
)

func HandleConnection(conn net.Conn) {
	var err error

	defer conn.Close()

	s := bufio.NewScanner(conn)

	s.Scan()

	if err = s.Err(); err != nil {
		log.Printf("error scanning input: %v", err)
		sendInternalError(conn)
		return
	}

	input := s.Text()
	
	var res string
	if res, err = executeCommand(input); err != nil {
		log.Printf("error executing command: %v", err)
		sendInternalError(conn)
		return
	}

	if _, err = conn.Write([]byte(res)); err != nil {
		log.Printf("error sending response: %v", err)
	}
}

func sendInternalError(conn net.Conn) {
	conn.Write([]byte(formatError(string(constants.ErrorInternal))))
}