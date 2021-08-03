package client

import (
	"context"

	api "tictactoeweb/api"
)

type Ctx = context.Context

// Public

// Play starts the game by setting it up and running the game loop.
// It's a default bootstrapper.
// Other public functions are exposed for testing purposes.
func Play(ctx Ctx, dmn Domain, sr *api.StatusReply) error {
	// ctx, err := Setup(response)
	// if err != nil {
	// 	return err
	// }
	err := run(ctx, dmn.Game)
	if err != nil {
		return err
	}
	return nil
}

func run(ctx Ctx, g Game) (err error) {
	gam, more := g, get_status()
	for more {
		gam, more, err = Loop(gam)
		if err != nil {
			return err
		}
	}
	return nil
}
