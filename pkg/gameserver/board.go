package gameserver

import "strings"

type board [][]mark

const __ = "-"

func (b board) String() string {
	_, dump := reduce(b[:], []string{})
	return strings.Join(dump, "\n")
}

func reduce(board board, rows []string) ([][]string, []string) {
	if len(board) == 0 {
		return [][]string{}, rows
	}
	rowz := make([]string, len(rows))
	copy(rowz, rows)

	row := strings.Join(board[0][:], " ")
	rowz = append(rowz, row)
	return reduce(board[1:], rowz)
}

func blankBoard() board {
	return board{
		{__, __, __},
		{__, __, __},
		{__, __, __},
	}
}

func (b board) isFilled(c cell) bool {
	// WARN: possible out of range
	return b[c.row][c.col] != __
}

func (b board) hasEmpty() bool {
	for _, row := range b {
		for _, m := range row {
			if m == __ {
				return true
			}
		}
	}
	return false
}

func (b board) isWinner(m mark) bool {
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

func setCell(b board, c cell, m mark) board {
	// WARN: possible out of range
	b[c.row][c.col] = m
	return b
}
