package domain

import (
	"strings"
	. "tictactoeweb/internal"
)

// Public
type (
	Board struct {
		grid
	}
	Row  = []string
	Mark = string // "X" or "O" (or "No")
)

type grid = [Size][Size]Mark // Private

// Constants, Public
const (
	Size = 3
	Gap  = __
	X    = "X"
	O    = "O"

	X_x = "X_x"
)

const __ = "-" // Private

// Public

func BlankBoard() Board {
	return Board{
		grid: grid{
			{__, __, __},
			{__, __, __},
			{__, __, __},
		},
	}
}

func DeadBoard() Board {
	return Board{
		grid: grid{
			{X_x, X_x, X_x},
			{X_x, X_x, X_x},
			{X_x, X_x, X_x},
		},
	}
}

func NewBoard(gs ...grid) Board {
	if len(gs) == 1 {
		return Board{gs[0]}
	}
	return Board{BlankBoard().grid}
}

// Other
func (b Board) String() string {
	var dump []string
	for _, row := range b.grid {
		s := strings.Join(row[:], " ")
		dump = append(dump, s)
	}
	return strings.Join(dump, "\n")
}

// Checks

func (b Board) IsEmpty() Empty {
	grd := b.grid
	return b == Board{} || b == DeadBoard() ||
		len(grd) != Size ||
		len(grd[0]) != Size ||
		len(grd[1]) != Size ||
		len(grd[2]) != Size
}

func (b Board) HasEmpty() bool {
	for _, row := range b.grid {
		for _, m := range row {
			if m == Gap {
				return true
			}
		}
	}
	return false
}
