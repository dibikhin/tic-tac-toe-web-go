package game

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

func NewGame(id Id, gs ...string) Game {
	if len(gs) == 1 {
		return Game{
			id, NewBoard(NewId(), gs[0]), NoPlayer(), NoPlayer(),
		}
	}
	return Game{
		id, NewBoard(NewId()), NoPlayer(), NoPlayer(),
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

// Checks

func (g Game) IsReady() bool {
	return !g.Player1().IsEmpty() &&
		!g.Player2().IsEmpty() &&
		!g.Board().IsEmpty()
}
