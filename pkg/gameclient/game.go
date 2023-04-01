package gameclient

import (
	"fmt"

	"tictactoe/pkg/api"
)

type Game struct {
	status api.GameStatus

	player1   Player
	player2   Player
	playerWon Player

	board Board
}

type Board string

func printGame(g Game) {
	if g.status == api.GameStatus_NOT_STARTED {
		return
	}
	fmt.Println("\nGame:")
	fmt.Printf("\nPlayer 1: %v\n", g.player1.String())
	fmt.Printf("Player 2: %v\n", g.player2.String())
	fmt.Println()
	fmt.Println(g.board)
}
