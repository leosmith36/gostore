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
	for {
		var (
			input string
			err error
			conn net.Conn
			res string
		)

		fmt.Print("gostore> ")

		if input, err = bufio.NewReader(os.Stdin).ReadString('\n'); err != nil {
			fmt.Printf("ERROR failed to read input: %v\n", err)
		}

		input = strings.TrimSpace(input)
		if input == "" {
			continue
		}
		
		if conn, err = net.Dial("tcp", ":8080"); err != nil {
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