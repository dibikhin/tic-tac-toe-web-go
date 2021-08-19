package client

import (
	api "tictactoeweb/api"

	. "tictactoeweb/internal"
)

// Public

// Play starts the game by setting it up and running the game loop.
// It's a default bootstrapper.
// Other public functions are exposed for testing purposes.
func Play(ctx Ctx, sr *api.StatusReply) error {
	// ctx, err := Setup()
	// if err != nil {
	// 	return err
	// }
	err := run(ctx)
	if err != nil {
		return err
	}
	return nil
}

func run(ctx Ctx) (err error) {
	game, more := game{}, Yes //get_status()
	for more {
		game, more, err = Loop(game)
		if err != nil {
			return err
		}
	}
	return nil
}
