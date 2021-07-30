// Package internal implements 3x3 Tic-tac-toe for 2 friends (cannot play with computer yet).
// Players choose their marks, put them, the Game checks winner or draw.
package internal

import (
	"errors"
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
	if !g.isReady() {
		return DeadGame(), false, ErrCouldNotStart()
	}
	gam, more := turn(g, g.player1)
	if !more {
		return gam, false, nil
	}
	gam, more = turn(gam, gam.player2)
	return gam, more, nil
}

// Private

func turn(g Game, playr player) (Game, bool) {
	c := inputLoop(g, playr)
	brd := setCell(g.board, c, playr.mark)
	brd.print()

	if brd.isWinner(playr.mark) {
		printWinner(playr)
		return setBoard(g, brd), false
	}
	if !brd.hasEmpty() {
		printDraw()
		return setBoard(g, brd), false
	}
	return setBoard(g, brd), true
}

func inputLoop(g Game, them player) cell {
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
