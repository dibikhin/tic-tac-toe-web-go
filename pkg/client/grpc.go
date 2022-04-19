package client

import (
	"log"
	"tictactoe/pkg/api"
	"tictactoe/pkg/config"

	"google.golang.org/grpc"
)

func Connect(cfg config.Config) (api.GameClient, func()) {
	log.Println("Connecting...")
	conn := grpcDial(cfg)
	client := api.NewGameClient(conn)
	log.Println("Connected")

	return client, func() {
		log.Println("Disconnected")
		conn.Close()
	}
}

func grpcDial(cfg config.Config) *grpc.ClientConn {
	conn, err := grpc.Dial(
		cfg.GameServer.Address,
		grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(cfg.GameServer.Timeout),
	)
	if err != nil {
		log.Fatalf("grpc: %v", err)
	}
	return conn
}
