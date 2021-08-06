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
	if !domain.Game.IsReady(g) {
		return domain.Game.DeadGame(), false, ErrCouldNotStart()
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
	brd := domain.Board.SetCell(g.Board().Id(), c, playr.Mark())
	brd.print()

	if brd.IsWinner(playr.Mark()) {
		printWinner(playr)
		return domain.Board.SetBoard(g, brd), false
	}
	if !brd.hasEmpty() {
		printDraw()
		return domain.Board.SetBoard(g, brd), false
	}
	return setBoard(g, brd), true
}

func inputLoop(g game, them Player) cell {
	prompt(them)
	for {
		turn := key(g.read())
		if !turn.isKey() {
			g.board.print()
			prompt(them)

			continue
		}
		cel := turn.toCell()
		if g.board.isFilled(cel) {
			g.board.print()
			prompt(them)

			continue
		}
		return cel
	}
}
