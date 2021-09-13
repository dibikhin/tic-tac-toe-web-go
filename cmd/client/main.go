package main

import (
	"context"
	"log"
	"sync"
	"tictactoeweb/internal/client"

	api "tictactoeweb/api"

	. "tictactoeweb/internal"
)

func main() {
	log.Print("App: starting...")

	log.Print("App: setting up input reader...")
	err := client.SetupReader()
	if err != nil {
		log.Fatalf("error: set up reader failed: %v", err)
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var wg sync.WaitGroup
	wg.Add(1)

	log.Print("Client: connecting to server...")
	teardown := WrapTeardown(cancel, wg.Done)
	conn, err := client.Start()
	if err != nil {
		teardown()
		log.Fatalf("error: start client failed: %v", err)
	}
	client.SetGameClient(api.NewGameClient(conn)) // Global

	OnExit(func() {
		client.Stop(conn, teardown)
	})
	defer client.Stop(conn, teardown)
	
	log.Print("App: running status loop...")
	err = client.RunStatusLoop(ctx)
	if err != nil {
		teardown()
		log.Fatalf("error: status loop broke: %v", err)
	}
	wg.Wait()
}
