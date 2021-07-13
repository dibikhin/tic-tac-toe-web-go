package main

import (
	"log"

	. "tictactoeweb/internal/server"
)

// clear && go run ./cmd/server/main.go
func main() {
	err := Start()
	if err != nil {
		log.Fatalf("Failed to start: %v", err)
	}
	log.Print("Started ok.")
}
