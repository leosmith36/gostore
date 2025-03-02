package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"net"
	"os"
	"strings"
)

func main() {
	var (
		err  error
		a    *args
		conn net.Conn
		res  string
	)

	if a, err = parseArgs(); err != nil {
		fmt.Print(err.Error())
		return
	}

	if conn, err = net.Dial("tcp", fmt.Sprintf("%s:%d", a.host, a.port)); err != nil {
		fmt.Println("ERROR failed to reach server")
		return
	}

	fmt.Fprintf(conn, "%s\n", "PING")

	if res, err = bufio.NewReader(conn).ReadString('\n'); err != nil && !errors.Is(err, io.EOF) {
		fmt.Println("ERROR failed to read response")
		return
	}

	if res != "PONG\n" {
		fmt.Println("ERROR unexpected response from server")
		return
	}

	for {
		var (
			input string
			res   string
			conn  net.Conn
		)

		fmt.Print("gostore> ")

		if input, err = bufio.NewReader(os.Stdin).ReadString('\n'); err != nil {
			fmt.Println("ERROR failed to read input")
		}

		input = strings.TrimSpace(input)
		if input == "" {
			continue
		}

		if conn, err = net.Dial("tcp", fmt.Sprintf("%s:%d", a.host, a.port)); err != nil {
			fmt.Println("ERROR failed to reach server")
			continue
		}

		fmt.Fprintf(conn, "%s\n", input)

		if res, err = bufio.NewReader(conn).ReadString('\n'); err != nil && !errors.Is(err, io.EOF) {
			fmt.Println("ERROR failed to read response")
			continue
		}

		fmt.Print(res)
	}
}
