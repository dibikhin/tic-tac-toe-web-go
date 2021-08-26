package game

import (
	"fmt"

	. "tictactoeweb/internal"
)

// Public
type (
	Board struct {
		id Id
		grid
	}
	Row  = struct{}
	Mark = string // "X" or "O" (or "No")
)

type grid string

// Constants, Public
const (
	Size = 3
	Gap  = __
	X    = "X"
	O    = "O"

	X_x = "X_x"
)

const __ = "-" // Private

func NewBoard(id Id, gs ...grid) Board {
	if len(gs) == 1 {
		return Board{id, gs[0]}
	}
	return Board{id: id}
}

func DeadMark() Mark {
	return X_x
}

// Other

func (b Board) String() string {
	return fmt.Sprintf("%v", b.grid)
}

// Props

func (b Board) Id() Id {
	return b.id
}

func (b Board) Grid() grid {
	return b.grid
}

// Checks

func (g grid) IsEmpty() Empty {
	return g == ""
}

func (b Board) IsEmpty() Empty {
	return b == Board{} || b.grid.IsEmpty()
}
