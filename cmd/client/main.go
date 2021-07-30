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
	log.Print("Starting client...")

	onExit(sayBye)

	teardown, cli, ctx, err := Start()
	if err != nil {
		log.Fatalf("could not get start: %v", err)
	}
	defer teardown()

	if err = StatusLoop(cli, ctx); err != nil {
		log.Fatalf("could not get status: %v", err)
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
}
