package domain

import (
	"fmt"
	"strings"

	irn "tictactoeweb/internal"
	g "tictactoeweb/internal/domain/game"
)

type (
	_Game  struct{}
	_Board struct{}

	reader = func() string
)

// Public

var Game = _Game{}
var Board = _Board{}

// Checks

func (_Game) IsReady(g g.Game) bool {
	return g.Reader() != nil &&
		!g.Player1().IsEmpty() &&
		!g.Player2().IsEmpty() &&
		!g.Board().IsEmpty()
}

// Constants

func (_Game) DeadGame() g.Game {
	return g.NewGame(g.DeadBoard())
}

func (_Game) NewGame() g.Game {
	return g.NewGame()
}

func (_Board) DeadBoard() g.Board {
	return g.NewBoard(g.DeadBoard())
}

func (_Board) NewBoard() g.Game {
	return g.NewBoard()
}

// Commands

func (_Game) SetPlayers(gm g.Game, p1, p2 g.Player) g.Game {
	// TODO: send to server
	return gm
}

func (_Board) SetBoard(gm g.Game, b g.Board) g.Game {
	// TODO: send to server
	return /*updated*/ gm
}

func (_Board) SetCell(boardId irn.Id, c g.Cell, m g.Mark) g.Board {
	// WARN: possible out of range
	// b[c.row][c.col] = m
	// TODO: send to server

	return g.Board{}
}

// Local

func (_Game) SetReader(gm g.Game, r reader) (g.Game, error) {
	return gm.SetReader(r, Game.DeadGame())
}

func (_Game) ChooseMarks(gm g.Game) (g.Player, g.Player, error) {
	fmt.Print("Press 'x' or 'o' to choose mark for Player 1: ")

	if gm.Reader() == nil {
		return g.DeadPlayer(), g.DeadPlayer(), irn.ErrNilReader()
	}
	m := gm.Reader()()
	p1, p2 := Game.ArrangePlayers(m)
	return p1, p2, nil
}

// Pure

func (_Game) ArrangePlayers(m g.Mark) (g.Player, g.Player) {
	if strings.ToLower(m) == "x" {
		return g.NewPlayer("X", 1), g.NewPlayer("O", 2)
	}
	return g.NewPlayer("O", 1), g.NewPlayer("X", 2)
}
