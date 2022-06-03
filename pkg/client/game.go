package client

import (
	"fmt"
	"tictactoe/pkg/api"
)

type game struct {
	status api.GameStatus

	player1   player
	player2   player
	playerWon player

	board string
}

func printGame(g game) {
	fmt.Println("\nGame:")
	fmt.Printf("\nPlayer 1: %v\n", g.player1.String())
	fmt.Printf("Player 2: %v\n", g.player2.String())
	fmt.Println()
	fmt.Println(g.board)
}
