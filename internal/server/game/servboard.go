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

// func BlankBoard() ServBoard {
// 	return NewServBoard{
// 		grid: grid{
// 			{__, __, __},
// 			{__, __, __},
// 			{__, __, __},
// 		},
// 	}
// }

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

func (b ServBoard) IsWinner(m Mark) bool {
	grd := b.grid
	// Horizontal
	h0 := grd[0][0] == m && grd[0][1] == m && grd[0][2] == m // 1 1 1 -> 7
	h1 := grd[1][0] == m && grd[1][1] == m && grd[1][2] == m // - - -
	h2 := grd[2][0] == m && grd[2][1] == m && grd[2][2] == m // - - -
	// Vertical
	v0 := grd[0][0] == m && grd[1][0] == m && grd[2][0] == m
	v1 := grd[0][1] == m && grd[1][1] == m && grd[2][1] == m
	v2 := grd[0][2] == m && grd[1][2] == m && grd[2][2] == m
	// Diagonal
	d0 := grd[0][0] == m && grd[1][1] == m && grd[2][2] == m
	d1 := grd[0][2] == m && grd[1][1] == m && grd[2][0] == m

	return h0 || h1 || h2 || v0 || v1 || v2 || d0 || d1
}
