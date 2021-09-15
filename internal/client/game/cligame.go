package game

import (
	. "tictactoeweb/internal"

	. "tictactoeweb/internal/domain/game"
)

type CliGame struct {
	Game
	player Player
}

func NewCliGame(id Id, gs ...string) CliGame {
	return CliGame{
		Game: NewGame(id, gs...),
	}
}

func (g CliGame) Board() CliBoard {
	return NewCliBoard(
		g.Game.Board().Id(),
		grid(g.Game.Board().Grid()),
	)
}

func (g CliGame) Player() Player {
	return g.player
}

func (g CliGame) SetPlayer(p Player) CliGame {
	g.player = p
	return g
}
