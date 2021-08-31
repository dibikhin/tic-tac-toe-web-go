package internal

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

// Other

type Ctx = context.Context

// App

func OnExit(wg sync.WaitGroup, done func()) {
	chos := make(chan os.Signal, 1)
	signal.Notify(chos, os.Interrupt, syscall.SIGTERM)
	go exit(wg, done, chos)
}

func SayBye() {
	fmt.Println()
	fmt.Println("Bye!")
	fmt.Println()
}

func exit(wg sync.WaitGroup, f func(), c chan os.Signal) {
	s := <-c
	log.Printf("App: got signal: %v. Exiting...", s)
	f()
	// log.Print("App: exited.")
	wg.Done()
	os.Exit(0)
}
