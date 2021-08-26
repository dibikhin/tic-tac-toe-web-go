package client

import (
	"tictactoeweb/api"

	. "tictactoeweb/internal"
)

// Public

// Globals

type (
	_App struct {
		reader Reader
	}
)

var (
	App    _App
	Client api.GameClient
)

// Props: Reader

func (a _App) Reader() Reader {
	return a.reader
}

func (a _App) SetReader(rdr Reader) error {
	if rdr == nil {
		return ErrNilReader()
	}
	a.reader = rdr
	return nil
}
