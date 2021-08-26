package game

import (
	. "tictactoeweb/internal"

	. "tictactoeweb/internal/domain/game"
)

type CliGame struct {
	Game

	reader Reader
}

func NewCliGame(id Id) CliGame {
	return CliGame{Game: NewGame(id)}
}

// Props:

func (g CliGame) Keys() []CliKey {
	return []CliKey{}
}

// Props: Reader

func (g CliGame) Reader() Reader {
	return g.reader
}

func (g CliGame) SetReader(rdr Reader, def CliGame) (CliGame, error) {
	if rdr == nil {
		return def, ErrNilReader()
	}
	g.reader = rdr
	return g, nil
}

// Checks

func (g CliGame) IsReady() bool {
	return g.Game.IsReady() &&
		g.reader != nil
}
