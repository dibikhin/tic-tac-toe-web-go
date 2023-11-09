package main

import (
	"fmt"
	"log"

	"tictactoe/app"
<<<<<<< Updated upstream
	"tictactoe/pkg/gameserver"
=======
>>>>>>> Stashed changes
)

func main() {
	log.Println("app: starting...")

	teardown := func() {}
<<<<<<< Updated upstream
	onExit := func() {
		teardown()
		log.Println("app: stopped")
		fmt.Println("\nBye!")
	}
	go app.WaitForExit(onExit)
	log.Println("app: started")

	cfg := app.LoadConfigFrom("./cmd/server/.env")
	lis := gameserver.Listen(cfg)
	s := gameserver.Make()

=======

	onExit := func() {
		teardown()
		log.Println("app: stopped")
		fmt.Println("\nBye!")
	}
	go app.WaitForExit(onExit)

	log.Println("app: started")

	cfg := app.LoadConfigFrom("./cmd/server/.env")
	lis := Listen(cfg)
	s := Make()

>>>>>>> Stashed changes
	teardown = func() {
		log.Println("app: gracefully stopping...")
		s.GracefulStop()
	}
<<<<<<< Updated upstream
	gameserver.Run(s, lis)
=======
	Run(s, lis)
>>>>>>> Stashed changes
}
