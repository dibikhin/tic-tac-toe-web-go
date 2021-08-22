package game

import (
	. "tictactoeweb/internal"
	domain "tictactoeweb/internal/domain/game"
)

type CliGame struct {
	domain.Game

	reader Reader
}

// Props:

func (g CliGame) Keys() []Key {
	return []Key{}
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
	return g.reader != nil &&
		g.IsReady()
}
