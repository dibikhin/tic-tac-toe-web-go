package domain

import (
	"strings"
	. "tictactoeweb/internal"
	. "tictactoeweb/internal/client/domain/game"
	. "tictactoeweb/internal/domain"
)

type (
	_Games  struct{} // see the `Games` var below
	_Boards struct{}

	reader = Reader
)

// Public

var Games = _Games{} // to call like `domain.Games.ArrangePlayers(m)`
var Boards = _Boards{}

func (_Games) ArrangePlayers(m Mark) (Game, error) {
	if strings.ToUpper(m) == X {
		// return NewPlayer(X, 1), NewPlayer(O, 2)
		return Game{}, nil
	}
	// return NewPlayer(O, 1), NewPlayer(X, 2)
	return Game{}, nil
}

func (_Games) Turn(boardId Id, trn Turn) (Game, error) {
	// WARN: possible out of range
	// b[c.row][c.col] = m
	return Game{}, nil
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
