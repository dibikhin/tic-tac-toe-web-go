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

type (
	Client = api.GameClient
	_Repo  = struct {
		SetById func()
	}
	Repo = struct {
		Client
		_Repo
	}
)

var _cli Client
var _repo Repo

// Repo.Start() ? TODO:

func Start() (Ctx, func(), error) {
	conn, err := grpc.Dial(cfg.Address, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(time.Second*8))
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	teardown := func() {
		conn.Close()
		cancel()
	}
	if err != nil {
		return nil, teardown, fmt.Errorf("did not connect: %w", err)
	}
	log.Print("Dialed address")

	// globals
	_cli = api.NewGameClient(conn)
	_repo = Repo{Client: _cli}

	log.Print("Connected client")

	return ctx, teardown, nil
}
