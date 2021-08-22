package client

import (
	"errors"
	"tictactoeweb/api"
	"tictactoeweb/internal/client/domain"
	"tictactoeweb/internal/domain/game"

	. "tictactoeweb/internal"

	. "tictactoeweb/internal/client/domain/game"
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
func Loop(g CliGame) (CliGame, again, error) {
	if !g.IsReady() {
		return domain.Games.MakeDead(), No, ErrCouldNotStart()
	}
	gam, more := turn(g.Player1(), g)
	if !more {
		return gam, No, nil
	}
	gam, more = turn(gam.Player2(), gam)
	return gam, more, nil
}

// Private

func turn(plr game.Player, gam CliGame) (CliGame, again) {
	trn := takeTurn(plr, gam)
	if trn == game.NoTurn() {
		return gam, No
	}
	gm, err := domain.Boards.Turn(gam.Board().Id(), trn)
	if err != nil {
		return gm, No
	}
	printOutcome(gm)
	return gm, Yes
}

func takeTurn(plr game.Player, gam CliGame) game.Turn {
	domain.Prompt(plr)
	more, t := Yes, game.NoTurn()
	for more {
		t, more = readTurn(plr, gam)
		if !more {
			return t
		}
	}
	return t
}

func printOutcome(gam CliGame) {
	domain.PrintBoard(gam.Board())

	switch o := gam.Outcome(); o {
	case api.Outcome_DRAW:
		domain.PrintDraw()
	case api.Outcome_WON:
		domain.PrintWinner(gam.Winner())
	}
}

func readTurn(plr game.Player, gam CliGame) (game.Turn, again) {
	read := gam.Reader()
	key := Key(read())
	if !key.IsIn(gam.Keys()) {
		domain.PrintBoard(gam.Board())
		domain.Prompt(plr)
		return game.NoTurn(), Yes
	}
	isFilled, err := domain.Boards.IsFilled(gam.Board().Id(), key)
	if err != nil {
		return game.NoTurn(), Yes
	}
	if isFilled {
		domain.PrintBoard(gam.Board())
		domain.Prompt(plr)
		return game.NoTurn(), Yes
	}
	return game.NewTurn(plr.Mark(), game.Key(key)), No
}
