package gameclient

import (
	"context"
	"fmt"
	"log"
	"time"

	"tictactoe/app"
	"tictactoe/pkg/api"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type service struct {
	c    api.GameClient
	cfg  app.Config
	read func() string
}

func NewService(c api.GameClient, cfg app.Config, r func() string) *service {
	return &service{c, cfg, r}
}

func (s *service) ReadPlayerName() name {
	return readName(s.read)
}

func (s *service) GetGame(ctx context.Context, nam name) game {
	for {
		r, err := s.c.GetGame(ctx, &api.GameRequest{PlayerName: nam})
		if err != nil {
			log.Printf("client: get game: %v", err)
			time.Sleep(s.cfg.Server.LoopDelay)
			continue
		}
		return makeGame(r)
	}
}

func makeGame(r *api.GameResponse) game {
	if r.Player1 == nil {
		return game{
			status:  r.Status,
			player1: player{},
			player2: player{},
			board:   r.Board,
		}
	}
	if r.Player2 == nil {
		return game{
			status:  r.Status,
			player1: player{mark: mark(r.Player1.Mark), name: r.Player1.Name},
			player2: player{},
			board:   r.Board,
		}
	}
	return game{
		status:    r.Status,
		player1:   player{mark: mark(r.Player1.Mark), name: r.Player1.Name},
		player2:   player{mark: mark(r.Player2.Mark), name: r.Player2.Name},
		playerWon: player{mark: mark(r.PlayerWon.Mark), name: r.PlayerWon.Name},
		board:     r.Board,
	}
}

func (s *service) StartGame(ctx context.Context, playerName name) {
	for {
		cmd := readCommand(s.read)
		if cmd == "p" {
			_, err := s.c.StartGame(ctx, &api.GameRequest{PlayerName: playerName})
			if err != nil {
				log.Printf("client: start game: %v", err)
				continue
			}
			break
		}
	}
}

func (s *service) Turn(ctx context.Context, p player) {
	for {
		t := readTurn(s.read, p.mark)
		_, err := s.c.Turn(ctx, &api.TurnRequest{PlayerName: p.name, Turn: t})
		if err != nil {
			log.Printf("client: turn: %v", err)
			status, _ := status.FromError(err)
			if status.Code() != codes.FailedPrecondition {
				continue
			}
		}
		break
	}
}

func readTurn(read func() string, mark mark) string {
	for {
		fmt.Printf("\nYour mark: %v. Press 1 to 9 (5 is center) and press ENTER: ", mark)
		turn := read()
		if turn == "" {
			continue
		}
		return turn
	}
}

func readCommand(read func() string) string {
	fmt.Println()
	for {
		fmt.Print("Type 'p' to play new game and press ENTER: ")
		cmd := read()
		if cmd == "" {
			continue
		}
		return cmd
	}
}

func readName(read func() string) name {
	fmt.Println()
	for {
		fmt.Print("What's your name? Type and press ENTER: ")
		name := read()
		if name == "" {
			continue
		}
		return name
	}
}
