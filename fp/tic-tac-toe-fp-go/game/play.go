// Package game is a bootstrapper only. It shows how to properly set up and run the game.
// See implementation in the internal package.
package game

import (
	. "tictactoe/internal"
)

// Public

// Play starts the game by setting it up and running the game loop.
// It's a default bootstrapper.
// Other public functions are exposed for testing purposes.
func Play() error {
	ctx, err := Setup()
	if err != nil {
		return err
	}
	more := true
	for more {
		ctx, more, err = Loop(ctx)
		if err != nil {
			return err
		}
	}
	return nil
}
