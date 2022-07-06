package main

import (
	"fmt"
	"log"

	"tictactoe/pkg/app"
	"tictactoe/pkg/server"
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

	cfg := app.LoadConfig()
	lis := server.StartListen(cfg)
	srv := server.MakeServer()

	teardown = func() {
		log.Println("app: gracefully stopping...")
		srv.GracefulStop()
	}
	server.RunServer(srv, lis)
}
