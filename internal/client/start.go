package client

import (
	"context"
	"log"
	cfg "tictactoeweb/configs"
	"time"

	"google.golang.org/grpc"

	api "tictactoeweb/api"
)

func Start() (*grpc.ClientConn, api.GameClient, context.Context, context.CancelFunc, error) {
	conn, err := grpc.Dial(cfg.Address, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(time.Second*8))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
		return nil, nil, nil, nil, err
	}
	log.Print("Dialed address")

	cli := api.NewGameClient(conn)
	log.Print("Connected client")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	return conn, cli, ctx, cancel, nil
}
