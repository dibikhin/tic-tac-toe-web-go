package data

import (
	. "tictactoeweb/internal"
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

func (_Games) GetById(boardId Id) (ServGame, error) {
	g, ok := DB().Games[boardId]
	if !ok {
		return DeadGame(), nil // TODO: errNotFound
	}
	return g, nil
}

func (_Boards) GetById(boardId Id) (ServBoard, error) {
	b, ok := DB().Boards[boardId]
	if !ok {
		return DeadBoard(), nil // TODO: errNotFound
	}
	return b, nil
}
