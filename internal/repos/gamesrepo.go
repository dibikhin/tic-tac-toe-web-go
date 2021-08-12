package internal

import (
	irn "tictactoeweb/internal"
	. "tictactoeweb/internal/domain/game"
)

type GamesRepo interface {
}

type State int

const (
	Unknown State = iota
	Connected
	Disconnected
)

type GamesDb struct {
	games []Game
	state State
}

func (GamesDb) Connect() GamesDb {
	return GamesDb{state: Connected}
}

func (GamesDb) Disconnect() GamesDb {
	return GamesDb{state: Disconnected}
}

func (db GamesDb) GetById(id irn.Id) Game {
	return NewGame(id)
}

func (db GamesDb) SetById(id irn.Id) Game {
	return NewGame(id)
}

func (db GamesDb) UpdateById(id irn.Id) Game {
	return NewGame(id)
}

func (db GamesDb) DeleteById() bool {
	return true
}
