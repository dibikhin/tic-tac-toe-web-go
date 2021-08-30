package main

import (
	"context"
	"log"

	. "tictactoeweb/internal"
	. "tictactoeweb/internal/client"
)

func main() {
	_, cancel := context.WithCancel(context.Background())
	defer cancel()
	OnExit(cancel, SayBye)

	log.Print("Starting Client()...")

	err := SetupReader()
	if err != nil {
		log.Fatalf("error: start reader failed: %v", err)
	}

	log.Print("Connecting to server...")

	ctx, teardown, err := StartClient()
	if err != nil {
		log.Fatalf("error: start client failed: %v", err)
	}
	defer log.Print("Client exited.")
	defer teardown()

	log.Print("Running status loop...")

	err = RunStatusLoop(ctx)
	if err != nil {
		log.Fatalf("error: status loop broke: %v", err)
	}
}
