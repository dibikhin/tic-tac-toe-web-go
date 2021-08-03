package client

import (
	"context"
	"fmt"
	"log"
	cfg "tictactoeweb/configs"
	"time"

	"google.golang.org/grpc"

	api "tictactoeweb/api"
)

func Start() (func(), api.GameClient, context.Context, error) {
	conn, err := grpc.Dial(cfg.Address, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(time.Second*8))
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	teardown := func() {
		conn.Close()
		cancel()
	}
	if err != nil {
		return teardown, nil, nil, fmt.Errorf("did not connect: %w", err)
	}
	log.Print("Dialed address")
	cli := api.NewGameClient(conn)
	log.Print("Connected client")

	return teardown, cli, ctx, nil
}
