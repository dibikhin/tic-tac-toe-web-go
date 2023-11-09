package game

import (
	"errors"
	"strings"
)

type Board [][]Mark

const (
	Empty = Mark("-")
	__    = Empty
)

var ErrorOutOfRange = errors.New("cell out of range")

func (b Board) String() string {
	_, dump := reduce(b[:], []string{})
	return strings.Join(dump, "\n")
}

func blankBoard() Board {
	return Board{
		{__, __, __},
		{__, __, __},
		{__, __, __},
	}
}

func (b Board) IsFilled(c Cell) (bool, error) {
	if !c.isInRange(b) {
		return false, ErrorOutOfRange
	}
	return b[c.row][c.col] != __, nil
}

func (b Board) HasEmpty() bool {
	for _, row := range b {
		for _, m := range row {
			if m == __ {
				return true
			}
		}
	}
	return false
}

func (b Board) WithCell(c Cell, m Mark) (Board, error) {
	if !c.isInRange(b) {
		return blankBoard(), ErrorOutOfRange
	}
	b[c.row][c.col] = m
	return b, nil
}

// TODO: simplify?
func (b Board) IsWinner(m Mark) bool {
	// WARN: possible out of range

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

func toStringSlice(marks []Mark) []string {
	ss := make([]string, len(marks))
	for i, m := range marks {
		ss[i] = string(m)
	}
	return ss
}
