package main

import (
	"fmt"
	"log"

	"tictactoe/app"
<<<<<<< Updated upstream
	"tictactoe/pkg/api"
	"tictactoe/pkg/gameclient"
=======
	"tictactoe/client"
>>>>>>> Stashed changes
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
<<<<<<< Updated upstream
=======

>>>>>>> Stashed changes
	log.Println("app: started")

	fmt.Print("\nWelcome to web 3x3 Tic-tac-toe for 2 friends :)\n\n")

	cfg := app.LoadConfigFrom("./cmd/client/.env")
<<<<<<< Updated upstream
	var cl api.GameClient
	cl, teardown = gameclient.Connect(cfg)

	s := gameclient.NewService(cl, cfg, app.DefaultRead)
	gameclient.RunLoop(s, cfg)
=======
	cl, teardown := Connect(cfg)
	s := client.NewService(cl, cfg, app.DefaultRead)

	client.RunLoop(s, cfg)
>>>>>>> Stashed changes

	onExit()
}
