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

	reader = func() string
)

// Public

var Games = _Games{} // to call imported fns as `domain.Games.IsReady(g)`
var Boards = _Boards{}

// Checks

func (_Games) IsReady(g Game) bool {
	return g.Reader() != nil &&
		!g.Player1().IsEmpty() &&
		!g.Player2().IsEmpty() &&
		!Boards.IsEmpty(g.Board())
}

// Constants

func (_Games) Dead() Game {
	return NewGame(NewBoard, DeadBoard())
}

func (_Games) NewGame() Game {
	return NewGame(NewBoard)
}

func (_Boards) DeadBoard() Board {
	return NewBoard(DeadBoard())
}

func (_Boards) NewBoard() Game {
	return Game{}
}

// Commands

func (_Games) SetPlayers(g Game, p1, p2 Player) Game {
	// TODO: send to server
	return Game{}
}

func (_Boards) SetBoard(g Game, b Board) Game {
	// TODO: send to server
	return /*updated*/ Game{}
}

func (_Boards) SetCell(b Board, c Cell, m Mark) Board {
	// WARN: possible out of range
	// b[c.row][c.col] = m
	// TODO: send to server

	return Board{}
}

// Local Properties

func (_Games) SetReader(g Game, r reader) (Game, error) {
	return g.SetReader(r, Games.Dead())
}

// IO

func (_Games) ChooseMarks(g Game) (Player, Player, error) {
	fmt.Print("Press 'x' or 'o' to choose mark for Player 1: ")

	if g.Reader() == nil {
		return DeadPlayer(), DeadPlayer(), irn.ErrNilReader()
	}
	m := g.Reader()()
	p1, p2 := Games.ArrangePlayers(m)
	return p1, p2, nil
}

// Pure

func (_Games) ArrangePlayers(m Mark) (Player, Player) {
	if strings.ToLower(m) == "x" {
		return NewPlayer("X", 1), NewPlayer("O", 2)
	}
	return NewPlayer("O", 1), NewPlayer("X", 2)
}

func (_Boards) IsEmpty(b Board) bool {
	bb := b.Board()
	return b == Board{} ||
		b == DeadBoard() ||
		len(bb) != Size ||
		len(bb[0]) != Size ||
		len(bb[1]) != Size ||
		len(bb[2]) != Size
}

func (_Boards) IsFilled(b Board, c Cell) bool {
	// WARN: possible out of range
	return b.Board()[c.Row()][c.Col()] != Gap
}

func (_Boards) HasEmpty(b Board) bool {
	for _, row := range b.Board() {
		for _, m := range row {
			if m == Gap {
				return true
			}
		}
	}
	return false
}

func (_Boards) IsWinner(b Board, m Mark) bool {
	bb := b.Board()
	// Something better needed, too naive

	// Horizontal
	h0 := bb[0][0] == m && bb[0][1] == m && bb[0][2] == m
	h1 := bb[1][0] == m && bb[1][1] == m && bb[1][2] == m
	h2 := bb[2][0] == m && bb[2][1] == m && bb[2][2] == m

	// Vertical
	v0 := bb[0][0] == m && bb[1][0] == m && bb[2][0] == m
	v1 := bb[0][1] == m && bb[1][1] == m && bb[2][1] == m
	v2 := bb[0][2] == m && bb[1][2] == m && bb[2][2] == m

	// Diagonal
	d0 := bb[0][0] == m && bb[1][1] == m && bb[2][2] == m
	d1 := bb[0][2] == m && bb[1][1] == m && bb[2][0] == m

	return h0 || h1 || h2 || v0 || v1 || v2 || d0 || d1
}
