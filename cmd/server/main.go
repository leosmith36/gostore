package main

import (
	"context"
	"log"
	"lsmith/gostore/internal/server"
	"lsmith/gostore/internal/store"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func main() {
	var (
		err error
		ln net.Listener
		wg = new(sync.WaitGroup)
		st = store.NewStore()
	)

	ctx, cancel := context.WithCancel(context.Background())

	st.Start()

	if ln, err = net.Listen("tcp", ":8080"); err != nil {
		log.Fatalf("failed to start listener: %v", err)
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		
		for {
			var (
				err error
				conn net.Conn
			)

			if conn, err = ln.Accept(); err != nil {
				if ctx.Err() != nil {
					return
				}
				log.Printf("failed to accept connection: %v", err)
				continue
			}
	
			wg.Add(1)
			go func() {
				defer wg.Done()
				server.HandleConnection(ctx, conn, st)
			}()
		}
	}()

  log.Print("started application")

	sigch := make(chan os.Signal, 1)
	signal.Notify(sigch, syscall.SIGHUP, syscall.SIGTERM, syscall.SIGINT)

	<-sigch

	// cancel goroutines
	cancel()

	// close listener
	ln.Close()

	log.Print("shutting down application")

	// wait for goroutines to stop
	wg.Wait()

	// stop store
	st.Stop()

	log.Print("application shut down gracefully")
}