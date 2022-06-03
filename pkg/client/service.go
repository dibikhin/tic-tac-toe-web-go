package client

import (
	"context"
	"fmt"
	"log"
	"time"

	"tictactoe/pkg/api"
)

type gameService struct {
	c    api.GameClient
	read func() string
}

func NewGameService(c api.GameClient, r func() string) *gameService {
	return &gameService{c, r}
}

func (s *gameService) ReadName() string {
	return readName(s.read)
}

func (s *gameService) GetGame(ctx context.Context, name string) game {
	for {
		r, err := s.c.GetGame(ctx, &api.GameRequest{PlayerName: name})
		if err != nil {
			log.Printf("client: %v", err)
			time.Sleep(time.Second)
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

func (s *gameService) StartGame(ctx context.Context, playerName string) {
	for {
		cmd := readCommand(s.read, playerName)
		if cmd == "p" {
			_, err := s.c.StartGame(ctx, &api.GameRequest{PlayerName: playerName})
			if err != nil {
				log.Printf("client: %v", err)
				continue
			}
			break
		}
	}
}

func (s *gameService) Turn(ctx context.Context, p player) {
	for {
		t := readTurn(s.read, p)
		_, err := s.c.Turn(ctx, &api.TurnRequest{PlayerName: p.name, Turn: t})
		if err != nil {
			log.Printf("client: %v", err)
			continue
		}
		break
	}
}

func readTurn(read func() string, p player) string {
	for {
		fmt.Printf("\nYour mark: %v. Press 1 to 9 (5 is center) and press ENTER: ", p.mark)
		turn := read()
		if turn == "" {
			continue
		}
		return turn
	}
}

func readCommand(read func() string, name string) string {
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

func readName(read func() string) string {
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
