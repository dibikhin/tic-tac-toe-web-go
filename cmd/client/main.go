package main

import (
	"log"
	"sync"

	. "tictactoeweb/internal"
	. "tictactoeweb/internal/client"
)

func main() {
	log.Print("App: starting...")

	log.Print("App: setting up input reader...")
	err := SetupReader()
	if err != nil {
		log.Fatalf("error: start reader failed: %v", err)
	}
	wg := sync.WaitGroup{}

	log.Print("Client: connecting to server...")
	ctx, teardown, err := StartClient()

	done := WrapTeardown(teardown, wg.Done)
	wg.Add(1)
	OnExit(done)

	if err != nil {
		done()
		log.Fatalf("error: start client failed: %v", err)
	}
	log.Print("App: running status loop...")
	err = RunStatusLoop(ctx)
	if err != nil {
		done()
		log.Fatalf("error: status loop broke: %v", err)
	}
	wg.Wait()
}
