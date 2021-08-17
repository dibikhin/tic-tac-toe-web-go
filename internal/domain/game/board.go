package game

import (
	"strings"

	irn "tictactoeweb/internal"
)

type (
	Mark = string // "X" or "O" (or "No")
	Board struct {
		id   irn.Id
		grid grid
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

func (b Board) SetBoard(gr grid) Board {
	b.grid = gr
	return b
}

// Checks

// Party:Server
func (b Board) IsEmpty() bool {
	grd := b.grid
	return b == Board{} || b == Dead() || len(grd) != Size ||
		len(grd[0]) != Size ||
		len(grd[1]) != Size ||
		len(grd[2]) != Size
}

// Party:Server
func (b Board) IsFilled(c Cell) bool {
	// WARN: possible out of range
	return b.grid[c.Row()][c.Col()] != Gap
}

// Party:Server
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

// Party:Server
func (b Board) IsWinner(m Mark) bool {
	gg := b.grid
	// Horizontal
	h0 := gg[0][0] == m && gg[0][1] == m && gg[0][2] == m
	h1 := gg[1][0] == m && gg[1][1] == m && gg[1][2] == m
	h2 := gg[2][0] == m && gg[2][1] == m && gg[2][2] == m
	// Vertical
	v0 := gg[0][0] == m && gg[1][0] == m && gg[2][0] == m
	v1 := gg[0][1] == m && gg[1][1] == m && gg[2][1] == m
	v2 := gg[0][2] == m && gg[1][2] == m && gg[2][2] == m
	// Diagonal
	d0 := gg[0][0] == m && gg[1][1] == m && gg[2][2] == m
	d1 := gg[0][2] == m && gg[1][1] == m && gg[2][0] == m

	return h0 || h1 || h2 || v0 || v1 || v2 || d0 || d1
}
