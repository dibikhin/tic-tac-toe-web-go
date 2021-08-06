package game

import (
	"strings"

	irn "tictactoeweb/internal"
)

type (
	Mark  = string // to avoid conversions
	board [_size][_size]Mark
	Board struct {
		id    irn.Id
		board board
	}
)

// Constants, Private

const (
	__    = "-"
	x_X   = "x_X"
	_size = 3
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
		irn.NewId(),
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

func NewBoard() Board {
	return Board{}
}

func (b Board) Id() irn.Id {
	return b.id
}

func (b Board) IsEmpty() bool {
	bb := b.board
	return b == Board{} ||
		b == DeadBoard() ||
		len(bb) != _size ||
		len(bb[0]) != _size ||
		len(bb[1]) != _size ||
		len(bb[2]) != _size
}

// Pure
func (b Board) isFilled(c Cell) bool {
	// WARN: possible out of range
	return b.board[c.row][c.col] != __
}

// Pure
func (b Board) hasEmpty() bool {
	for _, row := range b.board {
		for _, m := range row {
			if m == __ {
				return true
			}
		}
	}
	return false
}

// Pure
func (b Board) isWinner(m Mark) bool {
	bb := b.board
	// Something better needed, too naive

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

// No IO allowed in this file for SRP TODO:
