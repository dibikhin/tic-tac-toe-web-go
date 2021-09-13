package game

import (
	"strings"

	. "tictactoeweb/internal"
	. "tictactoeweb/internal/domain/game"
)

// Public

type (
	ServBoard struct {
		Board
		grid
	}
	// Row  [Size]string
	// Mark = string
)

type grid [Size][Size]Mark // Private

// Constants

const (
	__ = Gap
)

// Public

// Factorys

func BlankBoard() ServBoard {
	return ServBoard{
		grid: grid{
			{__, __, __},
			{__, __, __},
			{__, __, __},
		},
	}
}

// func DeadBoard() ServBoard {
// 	return NewServBoard{
// 		grid: grid{
// 			{X_x, X_x, X_x},
// 			{X_x, X_x, X_x},
// 			{X_x, X_x, X_x},
// 		},
// 	}
// }

// Other

func (b ServBoard) String() string {
	return b.Board.String()
}

func (g grid) String() string {
	var dump []string
	for _, row := range g {
		s := strings.Join(row[:], " ")
		dump = append(dump, s)
	}
	return strings.Join(dump, "\n")
}

// Props

// func (b ServBoard) SetBoard(gr grid) ServBoard {
// 	b.grid = gr
// 	return b
// }

func (b ServBoard) Grid() grid {
	return b.grid
}

// Checks, Validation

func (b ServBoard) IsFilled(c Cell) bool {
	// WARN: possible out of range
	return b.grid[c.Row()][c.Col()] != Gap
}

func (b ServBoard) HasEmpty() bool {
	for _, row := range b.grid {
		for _, m := range row {
			if m == Gap {
				return true
			}
		}
	}
	return false
}

func (g grid) IsEmpty() Empty {
	return len(g) != Size ||
		len(g[0]) != Size ||
		len(g[1]) != Size ||
		len(g[2]) != Size
}
