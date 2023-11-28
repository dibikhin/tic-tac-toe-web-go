package main

import (
	"fmt"
	"log"

	"tictactoe/app"
	"tictactoe/client"
)

func main() {
	log.Println("app: starting...")

	teardown := func() {}

	onExit := func() {
		teardown()

		log.Println("app: stopped")
		fmt.Println("\nBye!")
	}
	go app.WaitForExit(onExit)

	log.Println("app: started")

	fmt.Print("\nWelcome to web 3x3 Tic-tac-toe for 2 friends :)\n\n")

	cfg := app.LoadConfig("./cmd/client/.env")
	cl, teardown := Connect(cfg)
	s := client.NewGameService(cfg, cl, app.DefaultRead)

	client.RunGameLoop(s, cfg)

	onExit()
}
