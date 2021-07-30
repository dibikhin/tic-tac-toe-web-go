package game

import (
	"bufio"
	"os"
	"strings"

	gam "tictactoe/internal"
)

// User input strategy for stubbing in tests.
//
// NOTE: An interface is more idiomatic in this case. BUT it's overkill to define
// a type with constructor, an interface and its fake implementation in tests vs. this
// func, its impl and its fake impl in tests.
type reader = func() string

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
		return nil, gam.ErrNilReader()
	default:
		return rs[0], nil
	}
}
