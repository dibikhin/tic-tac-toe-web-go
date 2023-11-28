package main

import (
	"context"
	"log"

	"tictactoe/api"
	"tictactoe/app"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func Connect(cfg *app.Config) (api.GameClient, func()) {
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

func grpcDial(cfg *app.Config) *grpc.ClientConn {
	ctx, cancel := context.WithTimeout(context.TODO(), cfg.Server.Timeout)
	defer cancel()

	conn, err := grpc.DialContext(
		ctx,
		cfg.Server.Address,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		log.Fatalf("client: grpc dial: %v", err)
	}
	return conn
}
