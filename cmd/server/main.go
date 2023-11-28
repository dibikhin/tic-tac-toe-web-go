package main

import (
	"fmt"
	"log"

	"tictactoe/app"
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

	cfg := app.LoadConfig("./cmd/server/.env")
	lis := Listen(cfg.GameServer.Port)

	s, teardown := MakeServer()
	RunServer(s, lis)
}
