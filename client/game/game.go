package game

import (
	"fmt"

	"tictactoe/api"
)

type Game struct {
	Status api.GameStatus

	Player1   Player
	Player2   Player
	PlayerWon Player

	Board Board
}

type Board string

func MakeGame(r *api.GameResponse) Game {
	switch {
	case r.Player1 == nil:
		// A new game
		return Game{
			Status:  r.Status,
			Player1: emptyPlayer(),
			Player2: emptyPlayer(),
			Board:   Board(r.Board),
		}
	case r.Player2 == nil:
		// Player1 only
		return Game{
			Status:  r.Status,
			Player1: ToPlayer(r.Player1),
			Player2: emptyPlayer(),
			Board:   Board(r.Board),
		}
	default:
		// Two players
		return Game{
			Status:    r.Status,
			Player1:   ToPlayer(r.Player1),
			Player2:   ToPlayer(r.Player2),
			PlayerWon: ToPlayer(r.PlayerWon),
			Board:     Board(r.Board),
		}
	}
}

func emptyPlayer() Player {
	return Player{}
}

func PrintGame(g Game) {
	if g.Status == api.GameStatus_NOT_STARTED {
		return
	}
	fmt.Println("\nGame:")
	fmt.Printf("\nPlayer 1: %v\n", g.Player1.String())
	fmt.Printf("Player 2: %v\n", g.Player2.String())
	fmt.Println()
	fmt.Println(g.Board)
}
