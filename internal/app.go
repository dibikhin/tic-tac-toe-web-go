package internal

import (
	"context"
	"fmt"
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

	go func(f func(), c chan os.Signal) {
		defer wg.Done()
		<-c
		f()
		os.Exit(0)
	}(done, chos)
}

func SayBye() {
	fmt.Println()
	fmt.Println("Bye!")
	fmt.Println()
}
