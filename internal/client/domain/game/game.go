package game

import (
	. "tictactoeweb/internal"
	"tictactoeweb/internal/domain"
)

type Game struct {
	domain.Game

	reader Reader
}

func (g Game) Keys() []Key {
	return []Key{}
}

// Props: Reader

func (g Game) Reader() Reader {
	return g.reader
}

func (g Game) SetReader(rdr Reader, def Game) (Game, error) {
	if rdr == nil {
		return def, ErrNilReader()
	}
	g.reader = rdr
	return g, nil
}

// Checks

func (g Game) IsReady() bool {
	return g.reader != nil &&
		g.IsReady()
}
