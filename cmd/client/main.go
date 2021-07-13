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

	conn, cli, ctx, cancel, err := Start()
	if err != nil {
		log.Fatalf("could not get start: %v", err)
	}

	defer conn.Close()
	defer cancel()

	err = StatusLoop(cli, ctx)
	if err != nil {
		log.Fatalf("could not get status: %v", err)
	}
}

func onExit(done func()) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func(f func()) {
		<-c
		f()
		os.Exit(0)
	}(done)
}

func sayBye() {
	fmt.Println()
	fmt.Println("Bye!")
}
