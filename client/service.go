package client

import (
	"context"
	"fmt"
	"log"
	"time"

	"tictactoe/api"
	"tictactoe/app"
	"tictactoe/client/game"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type service struct {
	c    api.GameClient
	cfg  app.Config
	read func() string
}

func NewService(c api.GameClient, cfg app.Config, read func() string) *service {
	return &service{c, cfg, read}
}

func (s *service) ReadPlayerName() game.Name {
	return readName(s.read)
}

func (s *service) GetGame(ctx context.Context, name game.Name) game.Game {
	for {
		r, err := s.c.GetGame(ctx, &api.GameRequest{PlayerName: string(name)})
		if err != nil {
			log.Printf("client: get game: %v", err)

			time.Sleep(s.cfg.Server.LoopDelay)
			continue
		}
		return game.MakeGame(r)
	}
}

func (s *service) StartGame(ctx context.Context, playerName game.Name) error {
	for {
		cmd := readCommand(s.read)
		if cmd == "p" {
			_, err := s.c.StartGame(ctx, &api.GameRequest{PlayerName: string(playerName)})
			if err != nil {
				log.Printf("client: start game: %v", err)
				continue
			}
			break
		}
	}
	return nil
}

func (s *service) Turn(ctx context.Context, p game.Player) {
	for {
		t := readTurn(s.read, p.Mark)
		_, err := s.c.Turn(ctx, &api.TurnRequest{PlayerName: string(p.Name), Turn: t})
		if err != nil {
			log.Printf("client: turn: %v", err)

			status, _ := status.FromError(err)
			if status.Code() == codes.NotFound {
				return
			}
			// Retry transient errors
			continue
		}
		return
	}
}

func readTurn(read func() string, mark game.Mark) string {
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

func readName(read func() string) game.Name {
	fmt.Println()
	for {
		fmt.Print("What's your name? Type and press ENTER: ")
		name := read()
		if name == "" {
			continue
		}
		return game.Name(name)
	}
}
