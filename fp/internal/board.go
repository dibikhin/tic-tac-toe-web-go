package internal

import (
	"strings"
)

type (
	Board [_size][_size]mark

	mark = string // to avoid conversions
)

// Constants, Private

const (
	__    = "-"
	x_X   = "x_X"
	_size = 3
)

func _blankBoard() Board {
	return Board{
		{__, __, __},
		{__, __, __},
		{__, __, __}}
}

func _deadBoard() Board {
	return Board{
		{x_X, x_X, x_X},
		{x_X, x_X, x_X},
		{x_X, x_X, x_X}}
}

// Public

func (b Board) String() string {
	var dump []string
	for _, row := range b {
		s := strings.Join(row[:], " ")
		dump = append(dump, s)
	}
	return strings.Join(dump, "\n")
}

// Private

func setCell(b Board, c cell, m mark) Board {
	// WARN: possible out of range
	b[c.row][c.col] = m
	return b
}

// Pure
func (b Board) isEmpty() bool {
	return b == Board{} ||
		b == _deadBoard() ||

		len(b) != _size ||
		len(b[0]) != _size ||
		len(b[1]) != _size ||
		len(b[2]) != _size
}

// Pure
func (b Board) isFilled(c cell) bool {
	// WARN: possible out of range
	return b[c.row][c.col] != __
}

// Pure
func (b Board) hasEmpty() bool {
	for _, row := range b {
		for _, m := range row {
			if m == __ {
				return true
			}
		}
	}
	return false
}

// Pure
func (b Board) isWinner(m mark) bool {
	// Something better needed, too naive

	// Horizontal
	h0 := b[0][0] == m && b[0][1] == m && b[0][2] == m
	h1 := b[1][0] == m && b[1][1] == m && b[1][2] == m
	h2 := b[2][0] == m && b[2][1] == m && b[2][2] == m

	// Vertical
	v0 := b[0][0] == m && b[1][0] == m && b[2][0] == m
	v1 := b[0][1] == m && b[1][1] == m && b[2][1] == m
	v2 := b[0][2] == m && b[1][2] == m && b[2][2] == m

	// Diagonal
	d0 := b[0][0] == m && b[1][1] == m && b[2][2] == m
	d1 := b[0][2] == m && b[1][1] == m && b[2][0] == m

	return h0 || h1 || h2 || v0 || v1 || v2 || d0 || d1
}

// No IO allowed in this file for SRP
