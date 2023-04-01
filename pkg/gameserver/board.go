package gameserver

import "strings"

type Board [][]Mark

const __ = "-"

// TODO test
func (b Board) String() string {
	_, dump := reduce(b[:], []string{})
	return strings.Join(dump, "\n")
}

func reduce(board Board, rows []string) ([][]string, []string) {
	if len(board) == 0 {
		return [][]string{}, rows
	}
	rowz := make([]string, len(rows))
	copy(rowz, rows)

	ss := toStringSlice(board[0][:])
	row := strings.Join(ss, " ")
	rowz = append(rowz, row)
	return reduce(board[1:], rowz)
}

func toStringSlice(mark []Mark) []string {
	x := make([]string, len(mark))
	for i, m := range mark {
		x[i] = string(m)
	}
	return x
}

func blankBoard() Board {
	return Board{
		{__, __, __},
		{__, __, __},
		{__, __, __},
	}
}

// todo test
func (b Board) isFilled(c Cell) bool {
	// WARN: possible out of range
	return b[c.row][c.col] != __
}

// todo test
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

func (b Board) isWinner(m Mark) bool {
	// Horizontal
	// h0 := b[0][0] == m && b[0][1] == m && b[0][2] == m
	// h1 := b[1][0] == m && b[1][1] == m && b[1][2] == m
	// h2 := b[2][0] == m && b[2][1] == m && b[2][2] == m

	for i := 0; i < len(b); i++ {
		for j := 0; j < len(b); j++ {
			if b[j][i] != m {
				return false
			}
		}
	}

	// Vertical
	// v0 := b[0][0] == m && b[1][0] == m && b[2][0] == m
	// v1 := b[0][1] == m && b[1][1] == m && b[2][1] == m
	// v2 := b[0][2] == m && b[1][2] == m && b[2][2] == m

	for i := 0; i < len(b); i++ {
		for j := 0; j < len(b); j++ {
			if b[i][j] != m {
				return false
			}
		}
	}

	// TODO diagonal, clean

	// TODO fix & test

	// Diagonal
	// d0 := b[0][0] == m && b[1][1] == m && b[2][2] == m
	// d1 := b[0][2] == m && b[1][1] == m && b[2][0] == m

	// return h0 || h1 || h2 || v0 || v1 || v2 ||
	// return d0 || d1
	return true
}

// todo test
func setCell(b Board, c Cell, m Mark) Board {
	// WARN: possible out of range
	b[c.row][c.col] = m
	return b
}
