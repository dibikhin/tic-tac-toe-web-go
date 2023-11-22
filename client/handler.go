package client

import (
	"context"
	"fmt"
	"log"
	"time"

	"tictactoe/api"
	"tictactoe/app"
	"tictactoe/client/domain"
)

type Service interface {
	GetGame(context.Context, domain.PlayerName) domain.Game
	StartGame(context.Context, domain.PlayerName) error
	Turn(context.Context, domain.Player)

	ReadPlayerName() domain.PlayerName
}

func RunGameLoop(s Service, cfg *app.Config) {
	currentPlayer := s.ReadPlayerName()
	fmt.Printf("Hey %v!\n", currentPlayer)

	for {
		gam := s.GetGame(context.TODO(), currentPlayer)
		domain.PrintGame(gam)

		if gam.Status == api.GameStatus_SHUTDOWN_CLIENT {
			fmt.Println("\nGot shutdown command from server")
			return
		}
		if err := react(s, cfg, gam, currentPlayer); err != nil {
			log.Printf("react: %v", err)
		}
	}
}

func react(s Service, cfg *app.Config, gam domain.Game, playerName domain.PlayerName) error {
	delay := cfg.Server.LoopDelay

	switch gam.Status {

	case api.GameStatus_WON, api.GameStatus_DRAW,
		api.GameStatus_NOT_STARTED:

		fmt.Printf("\n%v %v\n", gam.Status, gam.PlayerWon.String())
		if err := s.StartGame(context.TODO(), playerName); err != nil {
			return fmt.Errorf("start game: player name %v %w", playerName, err)
		}
	case api.GameStatus_WAITING_P2_JOIN:
		startOrWait(s, delay, gam.Player1, playerName)

	case api.GameStatus_WAITING_P1_TO_TURN:
		turnOrWait(s, delay, gam.Player1, playerName)

	case api.GameStatus_WAITING_P2_TO_TURN:
		turnOrWait(s, delay, gam.Player2, playerName)

	default:
		fmt.Printf("\nUnknown status: %v\n", gam.Status)
		time.Sleep(delay)
	}
	return nil
}

func startOrWait(s Service, delay time.Duration, p domain.Player, playerName domain.PlayerName) {
	if p.Name != playerName {
		_ = s.StartGame(context.TODO(), playerName)
		return
	}
	fmt.Println("\nWaiting Player 2 to join...")
	log.Println()

	time.Sleep(delay)
}

func turnOrWait(s Service, delay time.Duration, p domain.Player, playerName domain.PlayerName) {
	if p.Name == playerName {
		s.Turn(context.TODO(), p)
		return
	}
	fmt.Printf("\nWaiting Player %v to turn...\n", playerName)
	log.Println()

	time.Sleep(delay)
}
