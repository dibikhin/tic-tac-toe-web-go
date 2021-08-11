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
	X_x  = "X_x"
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

func Dead() Board {
	return Board{
		X_x,
		board{
			{X_x, X_x, X_x},
			{X_x, X_x, X_x},
			{X_x, X_x, X_x},
		},
	}
}

// Other, Public

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
			bs[0].id, bs[0].board,
		}
	}
	return Board{
		irn.NewId(), BlankBoard().board,
	}
}

// Props

func (b Board) Id() irn.Id {
	return b.id
}

func (b Board) Board() board {
	return b.board
}

// Checks

func (b Board) IsEmpty() bool {
	bb := b.Board()
	return b == Board{} || b == Dead() || len(bb) != Size ||
		len(bb[0]) != Size ||
		len(bb[1]) != Size ||
		len(bb[2]) != Size
}

func (b Board) IsFilled(c Cell) bool {
	// WARN: possible out of range
	return b.Board()[c.Row()][c.Col()] != Gap
}

func (b Board) HasEmpty() bool {
	for _, row := range b.Board() {
		for _, m := range row {
			if m == Gap {
				return true
			}
		}
	}
	return false
}

func (b Board) IsWinner(m Mark) bool {
	bb := b.Board()
	// Horizontal
	h0 := bb[0][0] == m && bb[0][1] == m && bb[0][2] == m
	h1 := bb[1][0] == m && bb[1][1] == m && bb[1][2] == m
	h2 := bb[2][0] == m && bb[2][1] == m && bb[2][2] == m
	// Vertical
	v0 := bb[0][0] == m && bb[1][0] == m && bb[2][0] == m
	v1 := bb[0][1] == m && bb[1][1] == m && bb[2][1] == m
	v2 := bb[0][2] == m && bb[1][2] == m && bb[2][2] == m
	// Diagonal
	d0 := bb[0][0] == m && bb[1][1] == m && bb[2][2] == m
	d1 := bb[0][2] == m && bb[1][1] == m && bb[2][0] == m

	return h0 || h1 || h2 || v0 || v1 || v2 || d0 || d1
}
