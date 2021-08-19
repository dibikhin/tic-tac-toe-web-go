package client

import (
	"errors"

	"tictactoeweb/api"
	
	. "tictactoeweb/internal"
	"tictactoeweb/internal/client/domain"
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
func Loop(g Game) (Game, again, error) {
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

func turn(plr Player, game Game) (Game, again) {
	trn := takeTurn(plr, game)
	if trn == NoTurn() {
		return game, No
	}
	gm := domain.Boards.Turn(game.Board().Id(), trn)
	printOutcome(gm)
	return gm, Yes
}

func takeTurn(plr Player, gam Game) Turn {
	domain.Prompt(plr)
	more, t := Yes, NoTurn()
	for more {
		t, more = readTurn(plr, gam)
		if !more {
			return t
		}
	}
	return t
}

func printOutcome(gam Game) {
	domain.PrintBoard(gam.Board())

	switch o := gam.Outcome(); o {
	case api.Outcome_DRAW:
		domain.PrintDraw()
	case api.Outcome_WON:
		domain.PrintWinner(gam.Winner())
	}
}

func readTurn(plr Player, game Game) (Turn, again) {
	read := game.Reader()
	key := Key(read())
	if !key.IsIn(game.Keys()) {
		domain.PrintBoard(game.Board())
		domain.Prompt(plr)
		return NoTurn(), Yes
	}
	if domain.Boards.IsFilled(game.Board().Id(), key) {
		domain.PrintBoard(game.Board())
		domain.Prompt(plr)
		return NoTurn(), Yes
	}
	return NewTurn(plr.Mark(), key), No
}
