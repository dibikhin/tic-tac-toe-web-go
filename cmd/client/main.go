package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"tictactoe/pkg/app"
	"tictactoe/pkg/client"
)

func main() {
	log.Println("app: starting...")

	cfg := app.LoadConfig()

	teardown := func() {}
	onExit := func() {
		teardown()
		log.Println("app: stopped")
		fmt.Println("\nBye!")
	}
	go waitForExit(onExit)
	log.Println("app: started")

	fmt.Print("\nWelcome to web 3x3 Tic-tac-toe for 2 friends :)\n\n")

	cl, teardown := client.Connect(cfg)
	s := client.NewGameService(cl, app.DefaultRead)
	client.RunGameLoop(s)

	onExit()
}

func waitForExit(onExit func()) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	sig := <-c

	log.Printf("app: got signal %v. Stopping...", sig)
	onExit()
	os.Exit(0)
}
