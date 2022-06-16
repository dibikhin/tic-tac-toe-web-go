package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"tictactoe/pkg/app"
	"tictactoe/pkg/server"
)

func main() {
	log.Println("app: starting...")

	cfg := app.LoadConfig()
	lis := server.StartListen(cfg)
	srv := server.MakeServer()

	go server.RunServer(srv, lis)
	log.Println("app: started")

	waitForExit()
	log.Println("app: gracefully stopping...")
	srv.GracefulStop()
	lis.Close()

	log.Println("app: stopped")
	fmt.Println("\nBye!")
}

func waitForExit() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	sig := <-c

	log.Printf("app: got signal %v. Stopping...", sig)
}
