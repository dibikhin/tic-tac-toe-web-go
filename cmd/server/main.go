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
	_, cancel := context.WithCancel(context.Background())
	defer cancel()

	wg := sync.WaitGroup{}
	wg.Add(1)
	OnExit(wg, func() {
		cancel()
		SayBye()
		// teardown()
	})

	// TODO: shutdown grpc
	_, err := Start()
	if err != nil {
		log.Fatalf("Failed to start: %v", err)
	}
	log.Print("Started ok.")
	wg.Wait()
}
