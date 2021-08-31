package main

import (
	"context"
	"log"
	"sync"

	. "tictactoeweb/internal"
	. "tictactoeweb/internal/server"
)

// clear && go run ./cmd/server/main.go
func main() {
	log.Print("Server: starting...")
	_, cancel := context.WithCancel(context.Background())
	defer cancel()

	wg := sync.WaitGroup{}
	teardown := func() {
		log.Print("Context: cancelling...")
		cancel()
		log.Print("Context: cancelled.")

		SayBye()
	}
	err := Start(wg, teardown)
	if err != nil {
		teardown()
		log.Fatalf("Server: failed to start: %v", err)
	}
	wg.Wait()
}
