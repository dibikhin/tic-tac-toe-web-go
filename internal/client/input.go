package client

import (
	"errors"

	"tictactoeweb/internal/domain"
	. "tictactoeweb/internal/domain/game"
)

// Constants, Private
// ErrCouldNotStart arises when `Loop()` is run without running `Setup()` first.
func ErrCouldNotStart() error {
	return errors.New("Game: couldn't start the Game loop, set up the Game first")
}

// Public

// Game Loop()

type again = bool

// Loop prompts players to take turns.
func Loop(g Game) (Game, again, error) {
	if !domain.Games.IsReady(g) {
		return domain.Games.Dead(), false, ErrCouldNotStart()
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

	if domain.Boards.IsWinner(brd, playr.Mark()) {
		printWinner(playr)
		return domain.Boards.SetBoard(g, brd), false
	}
	if !domain.Boards.HasEmpty(brd) {
		printDraw()
		return domain.Boards.SetBoard(g, brd), false
	}
	return domain.Boards.SetBoard(g, brd), true
}

func inputLoop(g Game, them Player) Cell {
	prompt(them)
	for {
		turn := Key(g.Reader()()) // TODO: .Read()
		if !turn.IsKey() {
			g.board.print()
			prompt(them)

			continue
		}
		cel := turn.ToCell()
		if domain.Boards.IsFilled(g.Board(), cel) {
			g.board.print()
			prompt(them)

			continue
		}
		return cel
	}
}
