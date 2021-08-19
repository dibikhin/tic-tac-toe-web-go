package game

import (
	. "tictactoeweb/internal"
)

type Game struct {
	id Id

	board Board

	player1 Player
	player2 Player

	reader Reader
}

func NewGame(gameId Id, bs ...Board) Game {
	if len(bs) == 1 {
		return Game{
			id: gameId, board: bs[0],
		}
	}
	return Game{
		id: gameId, board: NewBoard(NewId()),
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

func (g Game) Keys() []Key {
	return []Key{}
}

// Props: Reader

func (g Game) Reader() Reader {
	return g.reader
}

func (g Game) SetReader(rdr Reader, def Game) (Game, error) {
	if rdr == nil {
		return def, ErrNilReader()
	}
	g.reader = rdr
	return g, nil
}

// Checks

func (g Game) IsReady() bool {
	return g.reader != nil &&
		!g.player1.IsEmpty() &&
		!g.player2.IsEmpty() &&
		!g.board.IsEmpty()
}
