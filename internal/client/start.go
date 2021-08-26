package client

import (
	"context"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"

	api "tictactoeweb/api"
	cfg "tictactoeweb/configs"
)

type (
	Client = api.GameClient
)

// Globals

var _cli Client

func Start() (Ctx, func(), error) {
	conn, err := grpc.Dial(cfg.Address, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(time.Second*8))
	log.Print("Dialed address")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	teardown := func() {
		conn.Close()
		cancel()
	}
	if err != nil {
		return nil, teardown, fmt.Errorf("did not connect: %w", err)
	}
	_cli = api.NewGameClient(conn) // Global
	log.Print("Connected client")

	return ctx, teardown, nil
}
