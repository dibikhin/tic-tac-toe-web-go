package client

import (
	"context"
	"fmt"
	"log"
	"tictactoe/api"

	"google.golang.org/grpc"
)

// A stub for GameClient interface in pkg/api/tictactoe_grpc.pb.go

type apiClientStub struct {
	c int
	m map[int]*api.GameResponse
}

func (cs *apiClientStub) GetGame(
	ctx context.Context, in *api.GameRequest, opts ...grpc.CallOption,
) (*api.GameResponse, error) {
	fmt.Println()
	log.Println("GetGame()")

	cs.c++
	r := cs.m[cs.c]

	log.Println("Game status:", r.Status)
	return r, nil
}

func (cs *apiClientStub) StartGame(
	ctx context.Context, in *api.GameRequest, opts ...grpc.CallOption,
) (*api.EmptyResponse, error) {
	fmt.Println()
	log.Println("StartGame()")
	return &api.EmptyResponse{}, nil
}

func (cs *apiClientStub) Turn(
	ctx context.Context, in *api.TurnRequest, opts ...grpc.CallOption,
) (*api.EmptyResponse, error) {
	fmt.Println()
	log.Println("Turn()")
	return &api.EmptyResponse{}, nil
}
