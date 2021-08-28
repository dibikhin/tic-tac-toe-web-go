package game

import (
	. "tictactoeweb/internal"

	. "tictactoeweb/internal/domain/game"
)

type CliGame struct {
	Game
}

func NewCliGame(id Id) CliGame {
	return CliGame{Game: NewGame(id)}
}

// Props:

func (g CliGame) Keys() []CliKey {
	return []CliKey{}
}