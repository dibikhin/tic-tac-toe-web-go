package domain

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

func PrintGame(g Game) {
	if g.Status == api.GameStatus_NOT_STARTED {
		return
	}
	fmt.Printf("\nPlayer 1: %v\n", g.Player1.String())
	fmt.Printf("Player 2: %v\n", g.Player2.String())
	fmt.Println()
	fmt.Println(g.Board)
}

func MakeGame(r *api.GameResponse) Game {
	switch {
	case r.Player1 == nil:
		return newGame(r)
	case r.Player2 == nil:
		return onlyPlayer1Game(r)
	default:
		return twoPlayersGame(r)
	}
}

func newGame(r *api.GameResponse) Game {
	return Game{
		Status:  r.Status,
		Player1: emptyPlayer(),
		Player2: emptyPlayer(),
		Board:   Board(r.Board),
	}
}

func onlyPlayer1Game(r *api.GameResponse) Game {
	return Game{
		Status:  r.Status,
		Player1: ToPlayer(r.Player1),
		Player2: emptyPlayer(),
		Board:   Board(r.Board),
	}
}

func twoPlayersGame(r *api.GameResponse) Game {
	return Game{
		Status:    r.Status,
		Player1:   ToPlayer(r.Player1),
		Player2:   ToPlayer(r.Player2),
		PlayerWon: ToPlayer(r.PlayerWon),
		Board:     Board(r.Board),
	}
}

func emptyPlayer() Player {
	return Player{}
}
