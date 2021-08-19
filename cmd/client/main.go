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
	log.Print("Hi!")
	log.Print("This is 3x3 Tic-tac-toe for 2 friends :)")

	onExit(sayBye)

	log.Print("Trying to connect...")
	ctx, teardown, err := Start()
	defer teardown()

	if err != nil {
		log.Fatalf("error: start failed: %v", err)
	}
	if err = StatusLoop(ctx); err != nil {
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
}
