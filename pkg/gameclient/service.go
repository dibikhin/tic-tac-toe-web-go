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

func NewService(c api.GameClient, cfg app.Config, read func() string) *service {
	return &service{c, cfg, read}
}

func (s *service) ReadPlayerName() Name {
	return readName(s.read)
}

func (s *service) GetGame(ctx context.Context, name Name) Game {
	for {
		r, err := s.c.GetGame(ctx, &api.GameRequest{PlayerName: string(name)})
		if err != nil {
			log.Printf("client: get game: %v", err)

			time.Sleep(s.cfg.Server.LoopDelay)
			continue
		}
		return makeGame(r)
	}
}

func makeGame(r *api.GameResponse) Game {
	// New game
	if r.Player1 == nil {
		return Game{
			status:  r.Status,
			player1: Player{},
			player2: Player{},
			board:   Board(r.Board),
		}
	}
	// Player1 only
	if r.Player2 == nil {
		return Game{
			status:  r.Status,
			player1: NewPlayer(r.Player1),
			player2: Player{},
			board:   Board(r.Board),
		}
	}
	// Two players
	return Game{
		status:    r.Status,
		player1:   NewPlayer(r.Player1),
		player2:   NewPlayer(r.Player2),
		playerWon: NewPlayer(r.PlayerWon),
		board:     Board(r.Board),
	}
}

func (s *service) StartGame(ctx context.Context, playerName Name) {
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
}

func (s *service) Turn(ctx context.Context, p Player) {
	for {
		t := readTurn(s.read, p.mark)
		_, err := s.c.Turn(ctx, &api.TurnRequest{PlayerName: string(p.name), Turn: t})
		if err != nil {
			log.Printf("client: turn: %v", err)

			status, _ := status.FromError(err)
			if status.Code() == codes.NotFound {
				return
			}
			// For transient errors, retry
			continue
		}
		return
	}
}

func readTurn(read func() string, mark Mark) string {
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

func readName(read func() string) Name {
	fmt.Println()
	for {
		fmt.Print("What's your name? Type and press ENTER: ")
		name := read()
		if name == "" {
			continue
		}
		return Name(name)
	}
}
