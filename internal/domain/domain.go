package domain

import (
	"fmt"

	irn "tictactoeweb/internal"
	. "tictactoeweb/internal/domain/game"
)

type (
	_Games  struct{} // see the `Games` var below
	_Boards struct{}

	reader = irn.Reader
)

// Public

var Games = _Games{} // to call like `domain.Games.SetPlayers(g, p1, p2)`
var Boards = _Boards{}

// Factorys: Games etc.

func (_Games) Make() Game {
	return NewGame(irn.NewId())
}

func (_Games) MakeDead() Game {
	return NewGame(X_x, Dead())
}

// Commands: Remote (IO)

func (_Games) ArrangePlayers(m Mark) (Game, error) {
	// TODO: send to server
	return Game{}, nil
}

func (_Games) SetPlayers(g Game, p1, p2 Player) Game {
	// TODO: send to server
	return Game{}
}

func (_Boards) IsFilled(boardId irn.Id, cel Cell) bool {
	return true
}

func (_Boards) Turn(b Board, t Turn) Board {
	// WARN: possible out of range
	// b[c.row][c.col] = m
	// TODO: send to server

	return Board{}
}

// Commands: Local + IO

// ChooseMarks chooses players' marks as in a Google's TicTacToe doodle
func (_Games) ChooseMarks(g Game) (Mark, error) {
	if g.Reader() == nil {
		return "", irn.ErrNilReader()
	}
	read := g.Reader()
	return read(), nil
}

// IO

func PrintGame(g Game) {
	fmt.Println()

	fmt.Println(g.Player1())
	fmt.Println(g.Player2())

	PrintBoard(g.Board())
}

func PrintBoard(b Board) {
	// Explicit check for the interface
	var _ fmt.Stringer = b

	fmt.Println()
	fmt.Println()
	fmt.Println("Press 1 to 9 to mark an empty cell (5 is center), then press ENTER. Board:")
	fmt.Println()

	fmt.Println(b)
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
