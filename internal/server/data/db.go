package data

import (
	. "tictactoeweb/internal"
	. "tictactoeweb/internal/server/game"
)

type (
	_DB struct {
		Boards _BoardsDB
		Games  _GamesDB
	}
	_BoardsDB map[Id]ServBoard
	_GamesDB  map[Id]ServGame
)

// NOTE: Simplified injection. It's global anyway
var _db = &_DB{
	Games: _GamesDB{
		"123": BlankGame(),
	},
	Boards: _BoardsDB{
		"567": BlankBoard(),
	},
}

func DB() *_DB {
	return _db
}
