package client

import (
	"context"
	"fmt"
	"log"
	"time"

	"tictactoe/pkg/api"
)

type GameService interface {
	GetGame(context.Context, string) game
	StartGame(context.Context, string)
	Turn(context.Context, player)

	ReadName() string
}

func RunGameLoop(s GameService) {
	playerName := s.ReadName()
	fmt.Printf("Hey %v!\n", playerName)

	for {
		game := s.GetGame(context.TODO(), playerName)
		printGame(game)

		if game.status == api.GameStatus_SHUTDOWN_CLIENT {
			fmt.Println("\nGot shutdown command from server")
			return
		}
		processStatus(s, game, playerName)
	}
}

func processStatus(s GameService, game game, playerName string) {
	switch game.status {
	case api.GameStatus_NOT_STARTED:
		s.StartGame(context.TODO(), playerName)
	case api.GameStatus_WAITING_P2_JOIN:
		startOrSleep(s, game, playerName)
	case api.GameStatus_WAITING_P1_TO_TURN:
		turnOrWaitPlayer1(s, game, playerName)
	case api.GameStatus_WAITING_P2_TO_TURN:
		turnOrWaitPlayer2(s, game, playerName)
	case api.GameStatus_WON:
		fmt.Printf("\nPlayer %v won!\n", game.playerWon.String())
		s.StartGame(context.TODO(), playerName)
	case api.GameStatus_DRAW:
		fmt.Println("\nDraw")
		s.StartGame(context.TODO(), playerName)
	default:
		fmt.Printf("\nUnknown status: %v\n", game.status)
		time.Sleep(time.Second)
	}
}

func startOrSleep(s GameService, g game, playerName string) {
	if g.player1.name != playerName {
		s.StartGame(context.TODO(), playerName)
		return
	}
	fmt.Println("\nWaiting Player 2 to join...")
	log.Println()
	time.Sleep(time.Second)
}

func turnOrWaitPlayer1(s GameService, game game, playerName string) {
	if game.player1.name == playerName {
		s.Turn(context.TODO(), game.player1)
		return
	}
	fmt.Println("\nWaiting Player 1 to turn...")
	log.Println()
	time.Sleep(time.Second)
}

func turnOrWaitPlayer2(s GameService, game game, playerName string) {
	if game.player2.name == playerName {
		s.Turn(context.TODO(), game.player2)
		return
	}
	fmt.Println("\nWaiting Player 2 to turn...")
	log.Println()
	time.Sleep(time.Second)
}
