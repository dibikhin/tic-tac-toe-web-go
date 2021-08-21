package server

import (
	. "tictactoeweb/internal"
	"tictactoeweb/internal/domain"
)

type Game struct {
	domain.Game
}

// Checks

func (g Game) IsReady() Yes {
	return !g.Player1.IsEmpty() &&
		!g.player2.IsEmpty() &&
		!g.board.IsEmpty()
}

// func (b Board) SetBoard(gr grid) Board {
// 	b.grid = gr
// 	return b
// }

// func (b Board) IsFilled(c Cell) bool {
// 	// WARN: possible out of range
// 	return b.grid[c.Row()][c.Col()] != Gap
//

// // Party:Server
// func (b Board) IsWinner(m Mark) bool {
// 	grd := b.grid
// 	// Horizontal
// 	h0 := grd[0][0] == m && grd[0][1] == m && grd[0][2] == m // 1 1 1 -> 7
// 	h1 := grd[1][0] == m && grd[1][1] == m && grd[1][2] == m // - - -
// 	h2 := grd[2][0] == m && grd[2][1] == m && grd[2][2] == m // - - -
// 	// Vertical
// 	v0 := grd[0][0] == m && grd[1][0] == m && grd[2][0] == m
// 	v1 := grd[0][1] == m && grd[1][1] == m && grd[2][1] == m
// 	v2 := grd[0][2] == m && grd[1][2] == m && grd[2][2] == m
// 	// Diagonal
// 	d0 := grd[0][0] == m && grd[1][1] == m && grd[2][2] == m
// 	d1 := grd[0][2] == m && grd[1][1] == m && grd[2][0] == m

// 	return h0 || h1 || h2 || v0 || v1 || v2 || d0 || d1
// }
