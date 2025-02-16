package server

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"lsmith/gostore/internal/constants"
	"lsmith/gostore/internal/store"
	"net"
	"strings"
)

func HandleConnection(ctx context.Context, conn net.Conn, st *store.Store) {
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

	log.Printf("received command: %s", input)

	res := executeCommand(input, st)

	log.Printf("sending response: %s", res)

	if _, err = conn.Write([]byte(res)); err != nil {
		log.Printf("error sending response: %v", err)
	}
}

func executeCommand(input string, st *store.Store) (output string) {
	split := strings.Split(input, " ")

	if len(split) < 1 {
		return formatError(constants.ErrMissingCommand)
	}

	cmd := split[0]
	args := split[1:]

  switch (cmd) {
	case constants.InSet:
		return set(st, args...)
	case constants.InGet:
		return get(st, args...)
	case constants.InDel:
		return del(st, args...)
	}

	return formatError(fmt.Sprintf("unknown command: %s", cmd))
}

func sendInternalError(conn net.Conn) {
	conn.Write([]byte(formatError(constants.ErrInternal)))
}