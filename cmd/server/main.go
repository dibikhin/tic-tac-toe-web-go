package main

import (
	"context"
	"log"
	"sync"

	. "tictactoeweb/internal"
	"tictactoeweb/internal/server"
)

// clear && go run ./cmd/server/main.go
func main() {
	log.Print("App: starting...")

	_, cancel := context.WithCancel(context.Background())
	defer cancel()

	var wg sync.WaitGroup
	wg.Add(1)
	ServeWithTeardown(cancel, wg)
	wg.Wait()
}

func ServeWithTeardown(cancel context.CancelFunc, wg sync.WaitGroup) {
	teardown := WrapTeardown(cancel, wg.Done)
	err := Serve(teardown)
	if err != nil {
		wg.Add(1)
		teardown()
		log.Fatalf("App: failed to start: %v", err)
	}
}

func Serve(teardown func()) error {
	lis, srv, err := server.Prepare()
	if err != nil {
		return err
	}
	defer server.Stop(srv, teardown)
	OnExit(func() {
		server.Stop(srv, teardown)
	})
	return server.Start(srv, lis)
}
