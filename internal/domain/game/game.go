package game

import (
	irn "tictactoeweb/internal"
)

type Game struct {
	id irn.Id

	board   Board
	player1 Player
	player2 Player

	// Party:Client
	reader irn.Reader
}

func NewGame(gameId irn.Id, bs ...Board) Game {
	if len(bs) == 1 {
		return Game{
			id: gameId, board: bs[0],
		}
	}
	return Game{
		id: gameId, board: NewBoard(irn.NewId()),
	}
}

// Props

func (g Game) Id() irn.Id {
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

// Party:Client

// Props: Reader

func (g Game) Reader() irn.Reader {
	return g.reader
}

func (g Game) SetReader(rdr irn.Reader, def Game) (Game, error) {
	if rdr == nil {
		return def, irn.ErrNilReader()
	}
	g.reader = rdr
	return g, nil
}

// Checks

// Party:Client
func (g Game) IsReady() bool {
	return g.reader != nil &&
		!g.player1.IsEmpty() &&
		!g.player2.IsEmpty() &&
		!g.board.IsEmpty()
}
