package client

import (
	. "tictactoeweb/internal"
	. "tictactoeweb/internal/client/game"
)

// Public

// Play starts the game by setting it up and running the game loop.
// It's a default bootstrapper.
// Other public functions are exposed for testing purposes.
func Play(ctx Ctx) error {
	game, more, err := NewCliGame(), Yes, error(nil)
	for more {
		game, more, err = Loop(ctx, game)
		if err != nil {
			return err
		}
	}
	return nil
}
