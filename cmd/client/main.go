package main

import (
	"fmt"
	"log"

	"tictactoe/app"
	"tictactoe/pkg/api"
	"tictactoe/pkg/gameclient"
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

	cfg := app.LoadConfigFrom("./cmd/client/.env")
	var cl api.GameClient
	cl, teardown = gameclient.Connect(cfg)

	s := gameclient.NewService(cl, cfg, app.DefaultRead)
	gameclient.RunLoop(s, cfg)

	onExit()
}
