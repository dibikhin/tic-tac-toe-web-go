package client

import (
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"

	cfg "tictactoeweb/configs"
)

func Start() (*grpc.ClientConn, error) {
	log.Print("Client: connecting...")
	log.Print("gRPC: dialing address...")
	conn, err := grpc.Dial(cfg.Address, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(time.Second*8))

	// NOTE: should teardown on errors or ctrl-c while `grpc.Dial()`

	if err != nil {
		return nil, fmt.Errorf("StartClient(): did not connect: %w", err)
	}
	log.Print("gRPC: dialed address.")
	return conn, nil
}

func Stop(c *grpc.ClientConn) {
	if c == nil {
		log.Print("Client: connection == nil, ignored")
	} else {
		log.Print("Client: disconnecting...")
		c.Close()
		log.Print("Client: disconnected.")
	}
}
