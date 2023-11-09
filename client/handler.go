package client

import (
	"context"
	"fmt"
	"log"
	"time"

	"tictactoe/api"
	"tictactoe/app"
	"tictactoe/client/game"
)

type Service interface {
	GetGame(context.Context, game.Name) game.Game
	StartGame(context.Context, game.Name) error
	Turn(context.Context, game.Player)

	ReadPlayerName() game.Name
}

func RunLoop(s Service, cfg app.Config) {
	currentPlayer := s.ReadPlayerName()
	fmt.Printf("Hey %v!\n", currentPlayer)

	for {
		gam := s.GetGame(context.TODO(), currentPlayer)
		game.PrintGame(gam)

		if gam.Status == api.GameStatus_SHUTDOWN_CLIENT {
			fmt.Println("\nGot shutdown command from server")
			return
		}
		if err := react(s, cfg, gam, currentPlayer); err != nil {
			log.Printf("run loop: %v", err)
		}
	}
}

func react(s Service, cfg app.Config, gam game.Game, playerName game.Name) error {
	switch gam.Status {
	case api.GameStatus_NOT_STARTED:
		if err := s.StartGame(context.TODO(), playerName); err != nil {
			return fmt.Errorf("start game: %w", err)
		}
	case api.GameStatus_WAITING_P2_JOIN:
		startOrWait(s, cfg, gam.Player1, playerName)

	case api.GameStatus_WAITING_P1_TO_TURN:
		turnOrWait(s, cfg, gam.Player1, playerName, 1)

	case api.GameStatus_WAITING_P2_TO_TURN:
		turnOrWait(s, cfg, gam.Player2, playerName, 2)

	case api.GameStatus_WON:
		fmt.Printf("\nPlayer %v won!\n", gam.PlayerWon.String())
		s.StartGame(context.TODO(), playerName)

	case api.GameStatus_DRAW:
		fmt.Println("\nDraw")
		s.StartGame(context.TODO(), playerName)

	default:
		fmt.Printf("\nUnknown status: %v\n", gam.Status)
		time.Sleep(cfg.Server.LoopDelay)
	}
	return nil
}

func startOrWait(s Service, cfg app.Config, p game.Player, playerName game.Name) {
	if p.Name != playerName {
		s.StartGame(context.TODO(), playerName)
		return
	}
	fmt.Println("\nWaiting Player 2 to join...")
	log.Println()

	time.Sleep(cfg.Server.LoopDelay)
}

func turnOrWait(s Service, cfg app.Config, p game.Player, playerName game.Name, playerNum int) {
	if p.Name == playerName {
		s.Turn(context.TODO(), p)
		return
	}
	fmt.Printf("\nWaiting Player %v to turn...\n", playerNum)
	log.Println()

	time.Sleep(cfg.Server.LoopDelay)
}
