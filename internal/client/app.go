package client

import (
	"tictactoeweb/api"

	. "tictactoeweb/internal"
)

// App

// Globals, Private
// TODO: wrap with struct
var (
	_reader Reader
	_client api.GameClient
)

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

func GameClient() api.GameClient {
	return _client
}

func SetGameClient(c api.GameClient) error {
	// TODO: WARN nil check interface
	// if c == nil {
	// 	return err
	// }
	_client = c
	return nil
}

// to prevent the name collision
func GetReader() Reader {
	return _reader
}

func SetReader(rdr Reader) error {
	if rdr == nil {
		return ErrNilReader()
	}
	_reader = rdr
	return nil
}
