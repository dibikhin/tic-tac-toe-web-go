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

func SetupReader(rs ...Reader) error {
	alt, err := ExtractReader(rs...)
	if err != nil {
		return err
	}
	if alt != nil {
		return SetReader(alt)
	}
	return SetReader(DefaultReader)
}

func StartClient() (Ctx, func(), error) {
	log.Print("gRPC: dialing address...")
	conn, err := grpc.Dial(cfg.Address, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(time.Second*8))
	log.Print("gRPC: dialed address.")

	// NOTE: should teardown on errors or ctrl-c while `grpc.Dial()`

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	teardown := func() {
		if conn != nil {
			log.Print("Client: disconnecting...")
			conn.Close()
			log.Print("Client: disconnected.")
		}
		cancel()
	}
	if err != nil {
		return nil, teardown, fmt.Errorf("StartClient(): did not connect: %w", err)
	}
	SetClient(api.NewGameClient(conn)) // Global

	log.Print("Client: connected.")
	return ctx, teardown, nil
}
