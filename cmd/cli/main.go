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
		fmt.Printf("ERROR failed to reach server: %v\n", err)
		return
	}

	fmt.Fprintf(conn, "%s\n", "PING")

	if res, err = bufio.NewReader(conn).ReadString('\n'); err != nil && !errors.Is(err, io.EOF) {
		fmt.Printf("ERROR failed to read response: %v\n", err)
		return
	}

	if res != "PONG\n" {
		fmt.Println(res)
		fmt.Println("ERROR unexpected response from server")
		return
	}

	for {
		var (
			input string
			err   error
		)

		fmt.Print("gostore> ")

		if input, err = bufio.NewReader(os.Stdin).ReadString('\n'); err != nil {
			fmt.Printf("ERROR failed to read input: %v\n", err)
		}

		input = strings.TrimSpace(input)
		if input == "" {
			continue
		}

		if conn, err = net.Dial("tcp", fmt.Sprintf("%s:%d", a.host, a.port)); err != nil {
			fmt.Printf("ERROR failed to reach server: %v\n", err)
			continue
		}

		fmt.Fprintf(conn, "%s\n", input)

		if res, err = bufio.NewReader(conn).ReadString('\n'); err != nil && !errors.Is(err, io.EOF) {
			fmt.Printf("ERROR failed to read response: %v\n", err)
			continue
		}

		fmt.Print(res)
	}
}
