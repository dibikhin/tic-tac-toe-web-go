package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"tictactoe/pkg/app"
	"tictactoe/pkg/server"
)

func main() {
	log.Println("Starting...")
	log.Println("Started")

	cfg := app.LoadConfig()
	lis := server.StartListen(cfg)
	srv := server.MakeServer()

	go server.RunServer(srv, lis)

	waitForExit()
	srv.GracefulStop()
	lis.Close()

	log.Println("Stopped")
	log.Println("Bye!")
}

func waitForExit() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	sig := <-c

	log.Printf("Got signal: %v. Stopping...", sig)
}
