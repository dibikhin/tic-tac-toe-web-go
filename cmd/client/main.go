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
		log.Fatalf("App: cannot setup reader: %v", err)
	}
	log.Print("App: input reader ok.")

	var wg sync.WaitGroup
	ctx, cncl := context.WithCancel(context.Background())
	cancel := func() {
		wg.Add(1)
		cncl()
		wg.Done()
	}
	defer cancel()

	connectWithTeardown(cancel, ctx)
	wg.Wait()
}

func connectWithTeardown(cancel func(), ctx context.Context) {
	teardown := WrapCancel(cancel)

	log.Print("Client: connecting...")
	conn, err := client.Start()
	if err != nil {
		log.Fatalf("Client: failed to start: %v", err)
	}
	// Global
	client.SetGameClient(api.NewGameClient(conn))

	OnExit(func() {
		client.Stop(conn, teardown)
	})

	log.Print("App: running status loop...")
	err = client.RunStatusLoop(ctx)
	if err != nil {
		log.Fatalf("App: status loop broke: %v", err)
	}
}
