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
	gm, more := turn(ctx, g)
	return gm, more, nil
}

// Private

func turn(ctx Ctx, game CliGame) (CliGame, again) {
	trn := takeTurn(ctx, game.Player(), game.Board())
	if trn == NoTurn() {
		return game, No
	}
	gm, err := Domain.Boards.Turn(ctx, game.Board().Id(), trn)
	if err != nil {
		return gm, No
	}
	return gm, Yes
}

func takeTurn(ctx Ctx, plr Player, b CliBoard) Turn {
	Domain.Prompt(plr)
	more, t := Yes, NoTurn()
	for more {
		t, more = readTurn(ctx, plr, b)
		if !more {
			return t
		}
	}
	return t
}

func readTurn(ctx Ctx, plr Player, board CliBoard) (Turn, again) {
	read := GetReader() // checked on top
	key := read()
	turn := NewTurn(plr.Mark(), key)
	isFilled, err := Domain.Boards.IsFilled(ctx, board.Id(), key)
	if err != nil {
		return turn, Yes
	}
	if isFilled {
		Domain.PrintBoard(board)
		Domain.Prompt(plr)
		return turn, Yes
	}
	return turn, No
}
