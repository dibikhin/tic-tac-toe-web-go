package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	. "tictactoeweb/internal/client"
)

func main() {
	onExit(sayBye)

	log.Print("Starting client...")
	err := SetupReader()
	if err != nil {
		log.Fatalf("error: start reader failed: %v", err)
	}

	log.Print("Connecting to server...")
	ctx, teardown, err := StartClient()
	if err != nil {
		log.Fatalf("error: start client failed: %v", err)
	}
	defer teardown()

	log.Print("Running status loop...")
	err = RunStatusLoop(ctx)
	if err != nil {
		log.Fatalf("error: status loop failed: %v", err)
	}
}

func onExit(done func()) {
	cs := make(chan os.Signal, 1)
	signal.Notify(cs, os.Interrupt, syscall.SIGTERM)
	go func(f func(), c chan os.Signal) {
		<-c
		f()
		os.Exit(0)
	}(done, cs)
}

func sayBye() {
	fmt.Println()
	fmt.Println("Bye!")
	fmt.Println()
}
