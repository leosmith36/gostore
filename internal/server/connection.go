package server

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"lsmith/gostore/internal/constants"
	"lsmith/gostore/internal/types"
	"net"
	"time"
)

func HandleConnection(ctx context.Context, conn net.Conn, st types.KeyValueStore) {
	var err error

	defer conn.Close()

	conn.SetReadDeadline(time.Now().Add(30 * time.Second))
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

func executeCommand(input string, st types.KeyValueStore) (output string) {
	split := splitArgs(input)

	if len(split) < 1 {
		return formatError(constants.ErrorMissingCommand)
	}

	cmd := split[0]
	args := split[1:]

	switch cmd {
	case constants.InputSet:
		return set(st, args...)
	case constants.InputGet:
		return get(st, args...)
	case constants.InputDel:
		return del(st, args...)
	case constants.InputPing:
		return ping(args...)
	case constants.InputIncr:
		return incr(st, false, args...)
	case constants.InputDecr:
		return incr(st, true, args...)
	case constants.InputIncrBy:
		return incrBy(st, false, args...)
	case constants.InputDecrBy:
		return incrBy(st, true, args...)
	}

	return formatError(fmt.Sprintf("unknown command: %s", cmd))
}

func sendInternalError(conn net.Conn) {
	conn.Write([]byte(formatError(constants.ErrorInternal)))
}
