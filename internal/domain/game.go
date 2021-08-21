package domain

import (
	. "tictactoeweb/internal"
)

// Public

type Game struct {
	id Id

	board Board

	player1 Player
	player2 Player
}

func NewGame(gameId Id, bs ...Board) Game {
	if len(bs) == 1 {
		return Game{
			id: gameId, board: bs[0],
		}
	}
	return Game{
		id: gameId, board: NewBoard(),
	}
}

// Props

func (g Game) Id() Id {
	return g.id
}

func (g Game) Board() Board {
	return g.board
}

func (g Game) Player1() Player {
	return g.player1
}

func (g Game) Player2() Player {
	return g.player2
}


