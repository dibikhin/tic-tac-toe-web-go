package client

import (
	"context"
	"fmt"
	"log"
	"time"

	"tictactoe/api"
	"tictactoe/app"
	"tictactoe/client/domain"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type gameService struct {
	read func() string
	cl   api.GameClient
	cfg  *app.Config
}

func NewGameService(cfg *app.Config, cl api.GameClient, read func() string) *gameService {
	return &gameService{
		read, cl, cfg,
	}
}

func (s *gameService) ReadPlayerName() domain.PlayerName {
	return readName(s.read)
}

func (s *gameService) GetGame(ctx context.Context, name domain.PlayerName) domain.Game {
	for {
		resp, err := s.cl.GetGame(ctx, &api.GameRequest{
			PlayerName: string(name),
		})
		if err != nil {
			log.Printf("client: get game: %v", err)

			time.Sleep(s.cfg.Server.LoopDelay)
			continue
		}
		return domain.MakeGame(resp)
	}
}

func (s *gameService) StartGame(ctx context.Context, playerName domain.PlayerName) error {
	for {
		cmd := readCommand(s.read)
		if cmd == domain.StartGame {
			_, err := s.cl.StartGame(ctx, &api.GameRequest{
				PlayerName: string(playerName),
			})
			if err != nil {
				log.Printf("client: start game: %v", err)
				continue
			}
			break
		}
	}
	return nil
}

func (s *gameService) Turn(ctx context.Context, p domain.Player) {
	for {
		t := readTurn(s.read, p.Mark)

		_, err := s.cl.Turn(ctx, &api.TurnRequest{
			PlayerName: string(p.Name),
			Turn:       string(t),
		})
		if err != nil {
			log.Printf("client: turn: %v", err)

			status, _ := status.FromError(err)
			if status.Code() == codes.NotFound {
				return
			}
			// Retry for transient errors
			continue
		}
		return
	}
}

func readTurn(read func() string, mark domain.Mark) domain.Turn {
	for {
		fmt.Printf("\nYour mark: %v. Press 1 to 9 (5 is center) and press ENTER: ", mark)
		t := read()
		if t == "" {
			continue
		}
		return domain.Turn(t)
	}
}

func readCommand(read func() string) domain.Command {
	fmt.Println()
	for {
		fmt.Printf("Type '%v' to play new game and press ENTER: ", domain.StartGame)
		cmd := read()
		if cmd == "" {
			continue
		}
		return domain.Command(cmd)
	}
}

func readName(read func() string) domain.PlayerName {
	fmt.Println()
	for {
		fmt.Print("What's your name? Type and press ENTER: ")
		name := read()
		if name == "" {
			continue
		}
		return domain.PlayerName(name)
	}
}
