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
	log.Print("App: starting...")

	_, cancel := context.WithCancel(context.Background())
	defer cancel()

	wg := sync.WaitGroup{}
	wg.Add(1)
	teardown := WrapTeardown(cancel, wg.Done)
	err := Serve(teardown)
	if err != nil {
		wg.Add(1)
		teardown()
		log.Fatalf("App: failed to start: %v", err)
	}
	wg.Wait()
}
