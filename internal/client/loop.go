package client

import (
	"errors"
	// "tictactoeweb/api"

	. "tictactoeweb/internal"

	. "tictactoeweb/internal/client/game"
	. "tictactoeweb/internal/domain/game"
)

type again = bool

// Constants, Public
// ErrCouldNotStart arises when `Loop()` is run without running `Setup()` first.
func ErrCouldNotStart() error {
	return errors.New("Game: couldn't start the Game loop, set up the Game first")
}

// Public

// Game Loop()

// Loop prompts players to take turns.
func Loop(ctx Ctx, g CliGame) (CliGame, again, error) {
	if !g.IsReady() {
		return g, No, ErrCouldNotStart()
	}
	game, more := turn(ctx, g.Player1(), g)
	if !more {
		return game, No, nil
	}
	gm, more := turn(ctx, game.Player2(), game)
	return gm, more, nil
}

// Private

func turn(ctx Ctx, plr Player, game CliGame) (CliGame, again) {
	trn := takeTurn(ctx, plr, game)
	if trn == NoTurn() {
		return game, No
	}
	gm, err := Domain.Boards.Turn(ctx, game.Board().Id(), trn)
	if err != nil {
		return gm, No
	}
	return gm, Yes
}

func takeTurn(ctx Ctx, plr Player, g CliGame) Turn {
	Domain.Prompt(plr)
	more, t := Yes, NoTurn()
	for more {
		t, more = readTurn(ctx, plr, g)
		if !more {
			return t
		}
	}
	return t
}

func readTurn(ctx Ctx, plr Player, game CliGame) (Turn, again) {
	read := GetReader() // checked on top
	key := read()
	turn := NewTurn(plr.Mark(), key)
	isFilled, err := Domain.Boards.IsFilled(ctx, game.Board().Id(), key)
	if err != nil {
		return turn, Yes
	}
	if isFilled {
		Domain.PrintBoard(game.Board())
		Domain.Prompt(plr)
		return turn, Yes
	}
	return turn, No
}
