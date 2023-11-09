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

type gameService struct {
	read  func() string
	games api.GameClient
	cfg   app.Config
}

func NewGameService(cfg app.Config, cl api.GameClient, read func() string) *gameService {
	return &gameService{read, cl, cfg}
}

func (s *gameService) ReadPlayerName() game.Name {
	return readName(s.read)
}

func (s *gameService) GetGame(ctx context.Context, name game.Name) game.Game {
	for {
		r, err := s.games.GetGame(ctx, &api.GameRequest{PlayerName: string(name)})
		if err != nil {
			log.Printf("client: get game: %v", err)

			time.Sleep(s.cfg.Server.LoopDelay)
			continue
		}
		return game.MakeGame(r)
	}
}

func (s *gameService) StartGame(ctx context.Context, playerName game.Name) error {
	for {
		cmd := readCommand(s.read)
		if cmd == "p" {
			_, err := s.games.StartGame(ctx, &api.GameRequest{PlayerName: string(playerName)})
			if err != nil {
				log.Printf("client: start game: %v", err)
				continue
			}
			break
		}
	}
	return nil
}

func (s *gameService) Turn(ctx context.Context, p game.Player) {
	for {
		t := readTurn(s.read, p.Mark)
		_, err := s.games.Turn(ctx, &api.TurnRequest{PlayerName: string(p.Name), Turn: t})
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
