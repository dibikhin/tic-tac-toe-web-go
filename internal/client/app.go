package client

import (
	"tictactoeweb/api"

	. "tictactoeweb/internal"
)

// App

// Globals, Private
type _App struct {
	reader Reader
	client api.GameClient
}

var _app *_App

// Public

func SetupReader(rs ...Reader) error {
	alt, err := ExtractReader(rs...)
	if err != nil {
		return err
	}
	if alt != nil {
		return SetReader(alt)
	}
	return SetReader(DefaultReader)
}

// Props

func Api() api.GameClient {
	return _app.client
}

func SetApi(c api.GameClient) error {
	// TODO: WARN nil check interface
	// if c == nil {
	// 	return err
	// }
	_app.client = c
	return nil
}

// to prevent the name collision
func GetReader() Reader {
	return _app.reader
}

func SetReader(rdr Reader) error {
	if rdr == nil {
		return ErrNilReader()
	}
	_app.reader = rdr
	return nil
}
