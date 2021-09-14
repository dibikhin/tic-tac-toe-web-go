package server

import (
	"strings"

	. "tictactoeweb/internal"
	. "tictactoeweb/internal/domain/game"
	. "tictactoeweb/internal/server/data"
	. "tictactoeweb/internal/server/game"
)

type (
	_Domain struct {
		Games  _Games // see the `Games` var below
		Boards _Boards
	}
	_Games  struct{}
	_Boards struct{}
)

// Public

// Globals

var Domain = _Domain{} // to call like `Domain.Games.ArrangePlayers(m)`

// Remote

// Querys

func (_Games) GetState(gameId Id) (State, error) {
	g, err := Repos.Games.GetById(gameId)
	if err != nil {
		return UNDEFINED, err
	}
	return g.State(), nil
}

func (_Boards) IsFilled(boardId Id, key ServKey) (bool, error) {
	b, err := Repos.Boards.GetById(boardId)
	if err != nil {
		return No, err
	}
	isFilled := b.IsFilled(key.ToCell())
	return isFilled, nil
}

// Local
func (_Games) ExtractState(g ServGame) State {
	if g.Player1().IsEmpty() || g.Player2().IsEmpty() {
		return WAITING_FOR_MARK
	}
	// if Domain.Boards.IsWinner(g.Board(), g.Player1()) {
	// 	return GAME_OVER_P1_WON, p1
	// } else if Domain.Boards.IsWinner(g.Board(), g.Player2()) {
	// 	return GAME_OVER_P2_WON, p2
	// } else if g.Board().HasEmpty() {
	// 	return WAITING_FOR_TURN, p1 || p2
	// } else {
	// 	return GAME_OVER_DRAW
	// }
	return UNDEFINED
}

func (_Boards) IsWinner(b ServBoard, p Player) bool {
	grd := b.Grid()
	m := p.Mark()
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

// Commands

func (_Games) ArrangePlayers(m Mark) (ServGame, error) {
	if strings.ToUpper(m) == X {
		// return NewPlayer(X, 1), NewPlayer(O, 2)
		return ServGame{}, nil
	}
	// return NewPlayer(O, 1), NewPlayer(X, 2)
	return ServGame{}, nil
}

func (_Games) Turn(boardId Id, trn Turn) (ServGame, error) {
	// WARN: possible out of range
	// b[c.row][c.col] = m
	return ServGame{}, nil
}
