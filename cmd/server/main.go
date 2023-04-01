package main

import (
	"fmt"
	"log"

	"tictactoe/app"
	"tictactoe/pkg/gameserver"
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

	cfg := app.LoadConfigFrom("./cmd/server/.env")
	lis := gameserver.Listen(cfg)
	s := gameserver.Make()

	teardown = func() {
		log.Println("app: gracefully stopping...")
		s.GracefulStop()
	}
	gameserver.Run(s, lis)
}
