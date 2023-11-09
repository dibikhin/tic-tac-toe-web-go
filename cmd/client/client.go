package main

import (
	"log"

	"tictactoe/api"
	"tictactoe/app"

	"google.golang.org/grpc"
)

func Connect(cfg app.Config) (api.GameClient, func()) {
	log.Println("client: connecting...")
	conn := grpcDial(cfg)
	client := api.NewGameClient(conn)
	log.Println("client: connected")

	teardown := func() {
		log.Println("client: disconnecting...")
		conn.Close()
		log.Println("client: disconnected")
	}
	return client, teardown
}

func grpcDial(cfg app.Config) *grpc.ClientConn {
	conn, err := grpc.Dial(
		cfg.Server.Address,
		grpc.WithInsecure(),
		grpc.WithBlock(),
		grpc.WithTimeout(cfg.Server.Timeout),
	)
	if err != nil {
		log.Fatalf("client: grpc dial: %v", err)
	}
	return conn
}
