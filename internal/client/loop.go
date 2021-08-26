package client

import (
	"errors"
	"tictactoeweb/api"

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
func Loop(g CliGame) (CliGame, again, error) {
	if !g.IsReady() {
		return g, No, ErrCouldNotStart()
	}
	game, more := turn(g.Player1(), g)
	if !more {
		return game, No, nil
	}
	gm, more := turn(game.Player2(), game)
	return gm, more, nil
}

// Private

func turn(plr Player, game CliGame) (CliGame, again) {
	trn := takeTurn(plr, game)
	if trn == NoTurn() {
		return game, No
	}
	gm, err := Domain.Boards.Turn(game.Board().Id(), trn)
	if err != nil {
		return gm, No
	}
	printOutcome(gm)
	return gm, Yes
}

func takeTurn(plr Player, g CliGame) Turn {
	Domain.Prompt(plr)
	more, t := Yes, NoTurn()
	for more {
		t, more = readTurn(plr, g)
		if !more {
			return t
		}
	}
	return t
}

func printOutcome(game CliGame) {
	Domain.PrintBoard(game.Board())

	switch o := game.Outcome(); o {
	case api.Outcome_DRAW:
		Domain.PrintDraw()
	case api.Outcome_WON:
		Domain.PrintWinner(game.Winner())
	}
}

func readTurn(plr Player, game CliGame) (Turn, again) {
	read := game.Reader()
	key := CliKey(read())
	turn := NewTurn(plr.Mark(), Key(key))
	if !key.IsIn(game.Keys()) {
		Domain.PrintBoard(game.Board())
		Domain.Prompt(plr)
		return turn, Yes
	}
	isFilled, err := Domain.Boards.IsFilled(game.Board().Id(), key)
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
