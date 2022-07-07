package main

import (
	"fmt"
	"log"

	"tictactoe/pkg/app"
	"tictactoe/pkg/client"
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

	cfg := app.LoadConfig()
	cl, teardown := client.Connect(cfg)
	s := client.NewGameService(cl, app.DefaultRead)

	client.RunGameLoop(s)

	onExit()
}
