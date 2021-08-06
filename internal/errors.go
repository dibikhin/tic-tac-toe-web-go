package internal

import "errors"

// Constants

// ErrNilReader() arises when `Setup()` is run with nil reader.
func ErrNilReader() error {
	return errors.New("game: the reader is nil, use a non-nil reader or nothing for the default one while setting up")
}
