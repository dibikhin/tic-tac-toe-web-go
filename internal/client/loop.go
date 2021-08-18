package client

import (
	"errors"

	"tictactoeweb/api"
	"tictactoeweb/internal/domain"
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
func Loop(g Game) (Game, again, error) {
	if !g.IsReady() {
		return domain.Games.MakeDead(), false, ErrCouldNotStart()
	}
	gam, more := turn(g.Player1(), g)
	if !more {
		return gam, false, nil
	}
	gam, more = turn(gam.Player2(), gam)
	return gam, more, nil
}

// Private

func turn(plr Player, gam Game) (Game, bool) {
	t := takeTurn(plr, gam)
	if t == NoTurn() {
		return domain.Games.MakeDead(), false
	}
	brd := domain.Boards.Turn(gam.Board(), t)
	printOutcome(brd || remote_game)
	return remote_game
}

func takeTurn(plr Player, gam Game) Turn {
	domain.Prompt(plr)
	more, t := true, NoTurn()
	for more {
		t, more = readTurn(gam, plr)
		if !more {
			return t
		}
	}
	return NoTurn()
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

func readTurn(gam Game, plr Player) (Turn, again) {
	read := gam.Reader()
	turn := Key(read())
	if !turn.IsIn(gam.Keys()) {
		domain.PrintBoard(gam.Board())
		domain.Prompt(plr)
		return NoTurn(), true
	}
	cel := turn.ToCell()
	if domain.Boards.IsFilled(gam.Board().Id(), cel) {
		domain.PrintBoard(gam.Board())
		domain.Prompt(plr)
		return NoTurn(), true
	}
	return NewTurn(plr.Mark(), cel), false
}
