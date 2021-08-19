package domain

import (
	"fmt"

	. "tictactoeweb/internal"
	. "tictactoeweb/internal/domain/game"
)

type (
	_Games  struct{} // see the `Games` var below
	_Boards struct{}

	reader = Reader
)

// Public

var Games = _Games{} // to call like `domain.Games.ArrangePlayers(m)`
var Boards = _Boards{}

// Factorys: Games etc.

func (_Games) Make() Game {
	return NewGame(NewId())
}

func (_Games) MakeDead() Game {
	return NewGame(X_x, Dead())
}

// IO

// Commands: Remote

func (_Games) ArrangePlayers(m Mark) (Game, error) {
	return Game{}, nil
}

func (_Boards) Turn(boardId Id, trn Turn) Game {
	// WARN: possible out of range
	// b[c.row][c.col] = m
	return Game{}
}

// Querys: Remote
func (_Boards) IsFilled(boardId Id, key Key) bool {
	return Yes
}

// Commands: Local + IO

// ChooseMarks chooses players' marks as in a Google's TicTacToe doodle
func (_Games) ChooseMarks(g Game) (Mark, error) {
	if g.Reader() == nil {
		return "", ErrNilReader()
	}
	read := g.Reader()
	return read(), nil
}

// Local

func PrintGame(g Game) {
	fmt.Println()

	fmt.Println(g.Player1())
	fmt.Println(g.Player2())

	PrintBoard(g.Board())
}

func PrintBoard(brd Board) {
	// Explicit check for the interface
	var _ fmt.Stringer = brd

	fmt.Println()
	fmt.Println()
	fmt.Println("Press 1 to 9 to mark an empty cell (5 is center), then press ENTER. Board:")
	fmt.Println()

	fmt.Println(brd)
	fmt.Println()
}

func PrintWinner(p Player) {
	fmt.Printf("%v won!\n", p)
}

func PrintDraw() {
	fmt.Println("Draw!")
}

// Implicit check for `fmt.Stringer` impl
func Prompt(s fmt.Stringer) { // otherwise `type not defined in this package`
	fmt.Printf("%v, your turn: ", s)
}
