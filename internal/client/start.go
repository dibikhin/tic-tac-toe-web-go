package client

import (
	"context"
	"log"
	cfg "tictactoeweb/configs"
	"time"

	"google.golang.org/grpc"

	api "tictactoeweb/api"
)

func Start() (func(), api.GameClient, context.Context, error) {
	conn, err := grpc.Dial(cfg.Address, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(time.Second*8))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
		return func() { conn.Close() }, nil, nil, err
	}
	log.Print("Dialed address")

	cli := api.NewGameClient(conn)
	log.Print("Connected client")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	teardown := func() {
		conn.Close()
		cancel()
	}
	return teardown, cli, ctx, nil
}
