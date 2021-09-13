package data

import (
	. "tictactoeweb/internal"
	// . "tictactoeweb/internal/domain/game"
	. "tictactoeweb/internal/server/game"
)

type (
	_Repos struct {
		Games  _Games // see the `Games` var below
		Boards _Boards
	}
	_Games  struct{}
	_Boards struct{}
)

// Public

// Globals

var Repos = _Repos{} // to call like `Repos.Games.GetById(id)`

func (_Boards) GetById(boardId Id) (ServBoard, error) {
	return DB().Boards[boardId], nil
}
