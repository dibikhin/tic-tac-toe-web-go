package client

import (
	. "tictactoeweb/internal"
	. "tictactoeweb/internal/client/game"
	. "tictactoeweb/internal/domain/game"
)

// Public

// Play starts the game by setting it up and running the game loop.
// It's a default bootstrapper.
// Other public functions are exposed for testing purposes.
func Play(ctx Ctx, p Player, b Board) error {
	g := NewCliGame("asdf", string(b.Grid())) boardId
	game := g.SetPlayer(p)
	more, err := Yes, error(nil)
	for more {
		game, more, err = Loop(ctx, game)
		if err != nil {
			return err
		}
	}
	return nil
}
