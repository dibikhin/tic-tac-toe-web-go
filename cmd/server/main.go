package main

import (
	"context"
	"log"
	"sync"
	"tictactoeweb/internal/server"

	. "tictactoeweb/internal"
)

// clear && go run ./cmd/server/main.go
func main() {
	log.Print("App: starting...")

	_, cancel := context.WithCancel(context.Background())
	defer cancel()

	var wg sync.WaitGroup
	ServeWithTeardown(cancel, &wg)
	wg.Wait()
}

func ServeWithTeardown(cancel func(), wg *sync.WaitGroup) {
	teardown := WrapCancel(cancel)
	err := Serve(teardown, wg)
	if err != nil {
		log.Fatalf("App: failed to serve: %v", err)
	}
}

func Serve(teardown func(), wg *sync.WaitGroup) error {
	lis, srv, err := server.Prepare()
	if err != nil {
		return err
	}
	OnExit(func() {
		wg.Add(1)
		server.Stop(srv)
		teardown()
		wg.Done()
	})
	return server.Start(srv, lis)
}
