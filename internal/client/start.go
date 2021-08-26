package client

import (
	"context"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"

	api "tictactoeweb/api"
	cfg "tictactoeweb/configs"

	. "tictactoeweb/internal"
)

func StartClient() (Ctx, func(), error) {
	conn, err := grpc.Dial(cfg.Address, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(time.Second*8))
	log.Print("Dialed address")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	teardown := func() {
		conn.Close()
		cancel()
		log.Print("Disconnected client")
	}
	if err != nil {
		return nil, teardown, fmt.Errorf("did not connect: %w", err)
	}
	Client = api.NewGameClient(conn) // Global
	log.Print("Connected client")

	return ctx, teardown, nil
}

func SetupReader(rs ...Reader) error {
	alt, err := ExtractReader(rs...)
	if err != nil {
		return err
	}
	if alt != nil {
		return App.SetReader(alt)
	}
	return App.SetReader(DefaultReader)
}
