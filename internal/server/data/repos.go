package data

import (
	"fmt"
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

// Constants

func ErrNotFound(id Id) error {
	return fmt.Errorf("Repo: cannot find by id: %w", id)
}

// Repos

var Repos = _Repos{} // to call like `Repos.Games.GetById(id)`

func (_Games) GetById(id Id) (ServGame, error) {
	g, ok := DB().Games()[id]
	if !ok {
		return DeadGame(), ErrNotFound(id)
	}
	return g, nil
}

func (_Boards) GetById(id Id) (ServBoard, error) {
	b, ok := DB().Boards()[id]
	if !ok {
		return DeadBoard(), ErrNotFound(id)
	}
	return b, nil
}
