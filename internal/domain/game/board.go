package game

import (
	"strings"

	irn "tictactoeweb/internal"
)

type (
	Mark  = string // to avoid conversions
	board [Size][Size]Mark
	Board struct {
		id    irn.Id
		board board
	}
)

// Constants, Private

const (
	Gap  = __
	__   = "-"
	x_X  = "x_X"
	Size = 3
)

// Constants, Public

func Logo() Board {
	return Board{
		"logo",
		board{
			{"X", " ", "X"},
			{"O", "X", "O"},
			{"X", " ", "O"}},
	}
}

func BlankBoard() Board {
	return Board{
		__,
		board{
			{__, __, __},
			{__, __, __},
			{__, __, __},
		},
	}
}

func DeadBoard() Board {
	return Board{
		x_X,
		board{
			{x_X, x_X, x_X},
			{x_X, x_X, x_X},
			{x_X, x_X, x_X},
		},
	}
}

// Public

func (b Board) String() string {
	var dump []string
	for _, row := range b.board {
		s := strings.Join(row[:], " ")
		dump = append(dump, s)
	}
	return strings.Join(dump, "\n")
}

func NewBoard(bs ...Board) Board {
	if len(bs) == 1 {
		return Board{
			id: bs[0].id, board: bs[0].board}
	}
	return Board{
		id: irn.NewId(), board: BlankBoard().board,
	}
}

func (b Board) Id() irn.Id {
	return b.id
}

func (b Board) Board() board {
	return b.board
}
