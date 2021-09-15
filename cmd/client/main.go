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

	// Reader
	log.Print("App: setting up input reader...")
	err := client.SetupReader()
	if err != nil {
		log.Fatalf("App: cannot setup reader: %v", err)
	}
	log.Print("App: input reader ok.")

	// Connect
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var wg sync.WaitGroup
	ConnectWithTeardown(cancel, ctx, &wg)
	wg.Wait()
}

func ConnectWithTeardown(cancel func(), ctx context.Context, wg *sync.WaitGroup) {
	teardown := WrapCancel(cancel)
	err := Connect(teardown, ctx, wg)
	if err != nil {
		log.Fatalf("App: failed to connect: %v", err)
	}
}

func Connect(teardown func(), ctx context.Context, wg *sync.WaitGroup) error {
	conn, err := client.Start()
	if err != nil {
		return err
	}
	// Global
	client.SetApi(api.NewGameClient(conn))

	OnExit(func() {
		wg.Add(1)
		client.Stop(conn)
		teardown()
		wg.Done()
	})
	return client.RunStatusLoop(ctx)
}
