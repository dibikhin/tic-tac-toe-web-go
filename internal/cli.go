package internal

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
)

// CLI

// Puclic

func OnExit(exit func()) {
	chos := make(chan os.Signal, 1)
	signal.Notify(chos, os.Interrupt, syscall.SIGTERM)
	go doExit(exit, chos)
}

func WrapCancel(cancel func()) func() {
	return func() {
		log.Print("Context: cancelling...")
		cancel()
		log.Print("Context: cancelled.")
		SayBye()
	}
}

func SayBye() {
	fmt.Println()
	fmt.Println("Bye!")
	fmt.Println()
}

// Private

func doExit(exit func(), c chan os.Signal) {
	s := <-c
	log.Printf("App: got signal: %v. Exiting...", s)
	log.Print("App: tearing down...")
	exit()
	log.Print("App: teared down.")

	log.Print("App: exited.")
	os.Exit(0)
}
