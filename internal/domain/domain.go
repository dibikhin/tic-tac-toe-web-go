package domain

import (
	"fmt"
	"strings"

	irn "tictactoeweb/internal"
	. "tictactoeweb/internal/domain/game"
)

type (
	_Games  struct{} // see the `Games` var below
	_Boards struct{}

	reader = irn.Reader
	grid   = [Size][Size]Mark
)

// Public

var Games = _Games{} // to call like `domain.Games.SetBoard(g, b)`
var Boards = _Boards{}

// Constants
func Logo() Board {
	return NewBoard(
		"logo",
		grid{
			{X, " ", X},
			{O, X, O},
			{X, " ", O}},
	)
}

// Factorys: Games etc.

func (_Games) Make() Game {
	return NewGame(irn.NewId())
}

func (_Games) MakeDead() Game {
	return NewGame(X_x, Dead())
}

func (_Games) ArrangePlayers(m Mark) (Player, Player) {
	if strings.ToLower(m) == "x" {
		return NewPlayer(X, 1), NewPlayer(O, 2)
	}
	return NewPlayer(O, 1), NewPlayer(X, 2)
}

// Commands: Remote (IO)

func (_Games) SetPlayers(g Game, p1, p2 Player) Game {
	// TODO: send to server
	return Game{}
}

func (_Games) SetBoard(g Game, b Board) Game {
	// TODO: send to server
	return /*updated*/ Game{}
}

func (_Boards) SetCell(b Board, c Cell, m Mark) Board {
	// WARN: possible out of range
	// b[c.row][c.col] = m
	// TODO: send to server

	return Board{}
}

// Commands: Local (IO)

// ChooseMarks chooses players' marks as in a Google's TicTacToe doodle
func (_Games) ChooseMarks(g Game) (Player, Player, error) {
	fmt.Print("Press 'x' or 'o' to choose mark for Player 1: ")

	if g.Reader() == nil {
		return DeadPlayer(), DeadPlayer(), irn.ErrNilReader()
	}
	read := g.Reader()
	m := read()
	p1, p2 := Games.ArrangePlayers(m)
	return p1, p2, nil
}
