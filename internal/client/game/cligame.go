package game

import (
	. "tictactoeweb/internal"

	. "tictactoeweb/internal/domain/game"
)

type CliGame struct {
	Game
}

func NewCliGame(id Id, gs ...string) CliGame {
	return CliGame{
		NewGame(id, gs...),
	}
}
