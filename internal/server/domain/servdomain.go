package domain

import (
	"strings"
	. "tictactoeweb/internal"
	domain "tictactoeweb/internal/domain/game"
	. "tictactoeweb/internal/server/domain/game"
)

type (
	_Games  struct{} // see the `Games` var below
	_Boards struct{}

	reader = Reader
)

// Public

var (
	Games  = _Games{} // to call like `domain.Games.ArrangePlayers(m)`
	Boards = _Boards{}
)

func (_Games) ArrangePlayers(m Mark) (ServGame, error) {
	if strings.ToUpper(m) == domain.X {
		// return NewPlayer(X, 1), NewPlayer(O, 2)
		return ServGame{}, nil
	}
	// return NewPlayer(O, 1), NewPlayer(X, 2)
	return ServGame{}, nil
}

func (_Games) Turn(boardId Id, trn domain.Turn) (ServGame, error) {
	// WARN: possible out of range
	// b[c.row][c.col] = m
	return ServGame{}, nil
}

func (_Boards) IsFilled(boardId Id, key Key) (bool, error) {
	// // if brd.IsWinner(plr.Mark()) {
	// // 	domain.PrintWinner(plr)
	// // 	return domain.Games.SetBoard(gam, brd), false
	// // }
	// // if !brd.HasEmpty() {
	// // 	domain.PrintDraw()
	// // 	return domain.Games.SetBoard(gam, brd), false
	// // }
	// // return domain.Games.SetBoard(gam, brd), true
	return Yes, nil
}
