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

type grid struct{}

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
	return Board{grid{}}
}

func DeadBoard() Board {
	return Board{grid{}}
}

func NewBoard(gs ...grid) Board {
	if len(gs) == 1 {
		return Board{gs[0]}
	}
	return Board{BlankBoard().grid}
}

// Props

func (b Board) Id() Id {
	return b.id
}

// func (b Board) Grid() grid {
// 	return b.grid
// }

// Other

func (b Board) String() string {
	return fmt.Sprintf("%v", b.grid)
}

// Checks

func (g grid) IsEmpty() Empty {
	return g == grid{}
}

func (b Board) IsEmpty() Empty {
	return b == Board{} || b == DeadBoard() || b.grid.IsEmpty()
}
