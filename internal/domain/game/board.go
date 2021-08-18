package game

import (
	"strings"

	irn "tictactoeweb/internal"
)

type (
	Mark  = string // "X" or "O" (or "No")
	Board struct {
		id irn.Id
		grid
	}
)

type grid [Size][Size]Mark

// Constants, Public

const (
	Size = 3
	Gap  = __
	X    = "X"
	O    = "O"

	X_x = "X_x"
)

const __ = "-" // Private

func BlankBoard() Board {
	return Board{
		__,
		grid{
			{__, __, __},
			{__, __, __},
			{__, __, __},
		},
	}
}

func Dead() Board {
	return Board{
		X_x,
		grid{
			{X_x, X_x, X_x},
			{X_x, X_x, X_x},
			{X_x, X_x, X_x},
		},
	}
}

// Other, Public

// Party:Server
func (b Board) String() string {
	var dump []string
	for _, row := range b.grid {
		s := strings.Join(row[:], " ")
		dump = append(dump, s)
	}
	return strings.Join(dump, "\n")
}

func NewBoard(boardId irn.Id, gs ...grid) Board {
	if len(gs) == 1 {
		return Board{
			boardId, gs[0],
		}
	}
	return Board{
		irn.NewId(), BlankBoard().grid,
	}
}

// Props

func (b Board) Id() irn.Id {
	return b.id
}

func (b Board) IsEmpty() bool {
	return b.grid == ""
}
