package client

import (
	"context"
	"fmt"
	"log"
	"testing"

	grpc "google.golang.org/grpc"

	"tictactoe/pkg/api"
)

func TestRunGameLoop(t *testing.T) {
	tests := []struct{ name string }{{"Game loop"}}

	p1, p2 := makePlayers()
	cs := makeClientStub(p1, p2)

	c := -2
	m := makeKeySeq()
	s := NewGameService(cs, func() string {
		c++
		return m[c]
	})
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			RunGameLoop(s)
		})
	}
}

type clientStub struct {
	c int
	m map[int]*api.GameResponse
}

func (cs *clientStub) GetGame(
	ctx context.Context, in *api.GameRequest, opts ...grpc.CallOption,
) (*api.GameResponse, error) {
	fmt.Println()
	log.Println("GetGame()")
	cs.c++
	r := cs.m[cs.c]
	log.Println("Game status:", r.Status)
	return r, nil
}

func (cs *clientStub) StartGame(
	ctx context.Context, in *api.GameRequest, opts ...grpc.CallOption,
) (*api.EmptyResponse, error) {
	fmt.Println()
	log.Println("StartGame()")
	return &api.EmptyResponse{}, nil
}

func (cs *clientStub) Turn(
	ctx context.Context, in *api.TurnRequest, opts ...grpc.CallOption,
) (*api.EmptyResponse, error) {
	fmt.Println()
	log.Println("Turn()")
	return &api.EmptyResponse{}, nil
}

func makePlayers() (*api.Player, *api.Player) {
	p1 := &api.Player{
		Mark: "X",
		Name: "name1",
	}
	p2 := &api.Player{
		Mark: "O",
		Name: "name2",
	}
	return p1, p2
}

func makeKeySeq() map[int]string {
	m := map[int]string{
		-1: "",
		0:  "name1",
		1:  "p",
		2:  "0",
		3:  "",
		4:  "2",
		5:  "",
		6:  "5",
		7:  "",
		8:  "7",
		9:  "",
		10: "p",
	}
	return m
}

func makeClientStub(p1 *api.Player, p2 *api.Player) *clientStub {
	cs := &clientStub{
		c: -1,
		m: map[int]*api.GameResponse{
			0: {Status: api.GameStatus_NOT_STARTED},
			1: {
				Status:  api.GameStatus_WAITING_P2_JOIN,
				Player1: p1,
			},
			2: {
				Status:    api.GameStatus_WAITING_P1_TO_TURN,
				Player1:   p1,
				Player2:   p2,
				PlayerWon: &api.Player{},
				Board:     "- - -\n- - -\n- - -\n",
			},
			3: {
				Status:    api.GameStatus_WAITING_P2_TO_TURN,
				Player1:   p1,
				Player2:   p2,
				PlayerWon: &api.Player{},
				Board:     "X - -\n- - -\n- - -\n",
			},
			4: {
				Status:    api.GameStatus_WAITING_P1_TO_TURN,
				Player1:   p1,
				Player2:   p2,
				PlayerWon: &api.Player{},
				Board:     "X O -\n- - -\n- - -\n",
			},
			5: {
				Status:    api.GameStatus_WAITING_P2_TO_TURN,
				Player1:   p1,
				Player2:   p2,
				PlayerWon: &api.Player{},
				Board:     "X O X\n- - -\n- - -\n",
			},
			6: {
				Status:    api.GameStatus_WAITING_P1_TO_TURN,
				Player1:   p1,
				Player2:   p2,
				PlayerWon: &api.Player{},
				Board:     "X O X\nO - -\n- - -\n",
			},
			7: {
				Status:    api.GameStatus_WAITING_P2_TO_TURN,
				Player1:   p1,
				Player2:   p2,
				PlayerWon: &api.Player{},
				Board:     "X O X\nO X -\n- - -\n",
			},
			8: {
				Status:    api.GameStatus_WAITING_P1_TO_TURN,
				Player1:   p1,
				Player2:   p2,
				PlayerWon: &api.Player{},
				Board:     "X O X\nO X O\n- - -\n",
			},
			9: {
				Status:    api.GameStatus_WON,
				Player1:   p1,
				Player2:   p2,
				PlayerWon: p1,
				Board:     "X O X\nO X O\nX - -\n",
			},
			10: {
				Status:    api.GameStatus_SHUTDOWN_CLIENT,
				Player1:   p1,
				Player2:   p2,
				PlayerWon: &api.Player{},
				Board:     "",
			},
		},
	}
	return cs
}
