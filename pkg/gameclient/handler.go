package gameclient

import (
	"context"
	"fmt"
	"log"
	"time"

	"tictactoe/app"
	"tictactoe/pkg/api"
)

type Service interface {
	GetGame(context.Context, Name) Game
	StartGame(context.Context, Name)
	Turn(context.Context, Player)

	ReadPlayerName() Name
}

func RunLoop(s Service, cfg app.Config) {
	currentPlayer := s.ReadPlayerName()
	fmt.Printf("Hey %v!\n", currentPlayer)

	for {
		game := s.GetGame(context.TODO(), currentPlayer)
		printGame(game)

		if game.status == api.GameStatus_SHUTDOWN_CLIENT {
			fmt.Println("\nGot shutdown command from server")
			return
		}
		react(s, cfg, game, currentPlayer)
	}
}

func react(s Service, cfg app.Config, game Game, playerName Name) {
	switch game.status {
	case api.GameStatus_NOT_STARTED:
		s.StartGame(context.TODO(), playerName)

	case api.GameStatus_WAITING_P2_JOIN:
		startOrWait(s, cfg, game.player1, playerName)

	case api.GameStatus_WAITING_P1_TO_TURN:
		turnOrWait(s, cfg, game.player1, playerName, 1)

	case api.GameStatus_WAITING_P2_TO_TURN:
		turnOrWait(s, cfg, game.player2, playerName, 2)

	case api.GameStatus_WON:
		fmt.Printf("\nPlayer %v won!\n", game.playerWon.String())
		s.StartGame(context.TODO(), playerName)

	case api.GameStatus_DRAW:
		fmt.Println("\nDraw")
		s.StartGame(context.TODO(), playerName)

	default:
		fmt.Printf("\nUnknown status: %v\n", game.status)
		time.Sleep(cfg.Server.LoopDelay)
	}
}

func startOrWait(s Service, cfg app.Config, p Player, playerName Name) {
	if p.name != playerName {
		s.StartGame(context.TODO(), playerName)
		return
	}
	fmt.Println("\nWaiting Player 2 to join...")
	log.Println()

	time.Sleep(cfg.Server.LoopDelay)
}

func turnOrWait(s Service, cfg app.Config, p Player, playerName Name, playerNum int) {
	if p.name == playerName {
		s.Turn(context.TODO(), p)
		return
	}
	fmt.Printf("\nWaiting Player %v to turn...\n", playerNum)
	log.Println()

	time.Sleep(cfg.Server.LoopDelay)
}
