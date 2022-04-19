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
	log.Println("Starting...")
	cfg := app.LoadConfig()
	log.Println("Started")

	teardown := func() {}
	onExit := func() {
		teardown()
		log.Println("Stopped")
		fmt.Println("\nBye!")
	}
	go waitForExit(onExit)

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

	log.Printf("Got signal: %v. Stopping...", sig)
	onExit()
	os.Exit(0)
}
