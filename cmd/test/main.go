package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

func main() {
	var (
		conn net.Conn
		err error
		res string
	)

	if conn, err = net.Dial("tcp", ":8080"); err != nil {
		log.Fatalf("failed to dial: %v", err)
	}

	fmt.Fprintf(conn, "SET foo bar\n")

	if res, err = bufio.NewReader(conn).ReadString('\n'); err != nil && !errors.Is(err, io.EOF) {
		log.Fatalf("failed to read response: %v", err)
	}

	log.Printf("received response: %s", res)
	
	// if conn, err = net.Dial("tcp", ":8080"); err != nil {
	// 	log.Fatalf("failed to dial: %v", err)
	// }

	// fmt.Fprintf(conn, "DEL foo\n")

	// if res, err = bufio.NewReader(conn).ReadString('\n'); err != nil && !errors.Is(err, io.EOF) {
	// 	log.Fatalf("failed to read response: %v", err)
	// }

	// log.Printf("received response: %s", res)

	time.Sleep(1*time.Second)

	if conn, err = net.Dial("tcp", ":8080"); err != nil {
		log.Fatalf("failed to dial: %v", err)
	}

	fmt.Fprintf(conn, "GET foo\n")

	if res, err = bufio.NewReader(conn).ReadString('\n'); err != nil && !errors.Is(err, io.EOF) {
		log.Fatalf("failed to read response: %v", err)
	}

	log.Printf("received response: %s", res)
}