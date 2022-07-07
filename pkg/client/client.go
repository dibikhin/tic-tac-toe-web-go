package client

import (
	"log"

	"tictactoe/pkg/api"
	"tictactoe/pkg/config"

	"google.golang.org/grpc"
)

func Connect(cfg config.Config) (api.GameClient, func()) {
	log.Println("client: connecting...")
	conn := grpcDial(cfg)
	client := api.NewGameClient(conn)
	log.Println("client: connected")

	return client, func() {
		log.Println("client: disconnecting...")
		conn.Close()
		log.Println("client: disconnected")
	}
}

func grpcDial(cfg config.Config) *grpc.ClientConn {
	conn, err := grpc.Dial(
		cfg.GameServer.Address,
		grpc.WithInsecure(),
		grpc.WithBlock(),
		grpc.WithTimeout(cfg.GameServer.Timeout),
	)
	if err != nil {
		log.Fatalf("client: grpc dial: %v", err)
	}
	return conn
}
