package client

import (
	"tictactoeweb/api"

	. "tictactoeweb/internal"
)

// Globals

// Private
var (
	_reader Reader
	_client api.GameClient
)

// Public

// Props

func Client() api.GameClient {
	return _client
}

func SetClient(c api.GameClient) error {
	// TODO: WARN nil check interface
	if c == nil {
		return err
	}
	_client = c
	return nil
}

// Reader
// to prevent a name collision
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
