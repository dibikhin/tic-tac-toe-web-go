package data

import (
	. "tictactoeweb/internal"
	. "tictactoeweb/internal/server/game"
)

type (
	_DB struct {
		boards _BoardsDB
		games  _GamesDB
	}
	_BoardsDB map[Id]ServBoard
	_GamesDB  map[Id]ServGame
)

// NOTE: Simplified injection. It's global anyway
var _db = &_DB{
	games: _GamesDB{
		"123": BlankGame(),
	},
	boards: _BoardsDB{
		"567": BlankBoard(),
	},
}

func DB() *_DB {
	return _db
}

func (db *_DB) Games() _GamesDB {
	return db.games
}

func (db *_DB) Boards() _BoardsDB {
	return db.boards
}
