package internal

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
)

// Other

type Ctx = context.Context

// App

func OnExit(done func()) {
	chos := make(chan os.Signal, 1)
	signal.Notify(chos, os.Interrupt, syscall.SIGTERM)
	go handleExit(done, chos)
}

func WrapTeardown(cancel context.CancelFunc, done func()) func() {
	return func() {
		log.Print("Context: cancelling...")
		cancel()
		log.Print("Context: cancelled.")

		SayBye()
		done()
	}
}

func SayBye() {
	fmt.Println()
	fmt.Println("Bye!")
	fmt.Println()
}

func handleExit(exit func(), c chan os.Signal) {
	s := <-c
	log.Printf("App: got signal: %v. Exiting...", s)
	exit()
	log.Print("App: exited.")
	os.Exit(0)
}
