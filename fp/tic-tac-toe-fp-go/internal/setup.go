package internal

import (
	"bufio"
	"errors"
	"os"
	"strings"
)

// Constants, Private
var (
	// errNilReader arises when `Setup()` is run with nil reader.
	errNilReader = errors.New("game: the reader is nil, pass a non-nil reader or nothing for the default one")
)

// Setup initializes the game and helps players to choose marks.
// The param is a strategy for user input to be stubbed.
// One can pass nothing, the default reader is used in the case.
// Example:
// ctx, err := Setup()
// OR
// ctx, err := Setup(DefaultReader)
// OR
// ctx, err := Setup(yourReaderFunc)
func Setup(rs ...reader) (game, error) {
	alt, err := extractReader(rs)
	if err != nil {
		return _deadGame(), err
	}
	gam := makeGame(DefaultReader, alt)
	printLogo(gam.logo)

	defer gam.print()
	p1, p2 := gam.chooseMarks()
	return setPlayers(gam, p1, p2), nil
}

// DefaultReader gets player's input and returns it as a text.
// It's exposed as a default impl of the `reader` Strategy.
func DefaultReader() string {
	// NOTE: it's easier to create it in place on demand vs. to store
	// and to initialize it somewhere. The `NewScanner` is very cheap inside actually
	s := bufio.NewScanner(os.Stdin)
	s.Scan()
	return strings.TrimSpace(s.Text())

	// TODO: have to check and propagate _scanner.Err() ?
}

// Private

func extractReader(rs []reader) (reader, error) {
	switch {
	case len(rs) < 1:
		return nil, nil
	case rs[0] == nil:
		return nil, errNilReader
	default:
		return rs[0], nil
	}
}

// Factory
func makeGame(def, alt reader) game {
	gam := newGame()
	if alt != nil {
		return setReader(gam, alt)
	}
	return setReader(gam, def)
}
