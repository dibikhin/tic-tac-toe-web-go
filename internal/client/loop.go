package client

import (
	"errors"

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
	gam, more := turn(g, g.Player1())
	if !more {
		return gam, false, nil
	}
	gam, more = turn(gam, gam.Player2())
	return gam, more, nil
}

// Private

func turn(g Game, playr Player) (Game, bool) {
	c := inputLoop(g, playr)
	brd := domain.Boards.SetCell(g.Board(), c, playr.Mark())
	brd.print()

	if brd.IsWinner(playr.Mark()) {
		printWinner(playr)
		return domain.Games.SetBoard(g, brd), false
	}
	if !brd.HasEmpty() {
		printDraw()
		return domain.Games.SetBoard(g, brd), false
	}
	return domain.Games.SetBoard(g, brd), true
}

func inputLoop(g Game, them Player) Cell {
	prompt(them)
	for {
		read := g.Reader()
		turn := Key(read())
		if !turn.IsKey() {
			g.board.print()
			prompt(them)

			continue
		}
		cel := turn.ToCell()
		if g.Board().IsFilled(cel) {
			g.board.print()
			prompt(them)

			continue
		}
		return cel
	}
}
