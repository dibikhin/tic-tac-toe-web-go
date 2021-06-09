// Package internal implements 3x3 Tic-tac-toe for 2 friends (cannot play with computer yet).
// Players choose their marks, put them, the game checks winner or draw.
package internal

import (
	"errors"
	"fmt"
)

// Constants, Private
var (
	// errCouldNotStart arises when `Loop()` is run without running `Setup()` first.
	errCouldNotStart = errors.New("game: couldn't start the game loop, set up the game first")
)

// Public

// Game Loop()

type again = bool

// Loop prompts players to take turns.
func Loop(g game) (game, again, error) {
	if !g.isReady() {
		return _deadGame(), false, errCouldNotStart
	}
	gam, more := turn(g, g.player1)
	if !more {
		return gam, false, nil
	}
	gam, more = turn(gam, gam.player2)
	return gam, more, nil
}

// Private

func turn(g game, playr player) (game, bool) {
	c := inputLoop(g, playr)
	brd := setCell(g.board, c, playr.mark)
	brd.print()

	if brd.isWinner(playr.mark) {
		fmt.Printf("%v won!\n", playr)
		return setBoard(g, brd), false
	}
	if !brd.hasEmpty() {
		fmt.Println("Draw!")
		return setBoard(g, brd), false
	}
	return setBoard(g, brd), true
}

func inputLoop(g game, them player) cell {
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
