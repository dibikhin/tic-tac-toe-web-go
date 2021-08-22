package domain

import (
	"fmt"

	domain "tictactoeweb/internal/domain/game"

	. "tictactoeweb/internal"
	. "tictactoeweb/internal/client/domain/game"
)

type (
	_Games  struct{} // see the `Games` var below
	_Boards struct{}

	reader = Reader
)

// Public

var (
	Games  = _Games{} // to call like `domain.Games.ArrangePlayers(m)`
	Boards = _Boards{}
)

// Factorys: Games etc.

func (_Games) Make() CliGame {
	return NewGame(NewId())
}

func (_Games) MakeDead() CliGame {
	return NewGame(domain.X_x, DeadBoard())
}

// IO

// Commands: Remote

func (_Games) ArrangePlayers(m domain.Mark) (CliGame, error) {
	return CliGame{}, nil
}

func (_Boards) Turn(boardId Id, trn domain.Turn) (CliGame, error) {
	return CliGame{}, nil
}

// Querys: Remote
func (_Boards) IsFilled(boardId Id, key Key) (bool, error) {
	return Yes, nil
}

// Commands: Local + IO

// ChooseMarks chooses players' marks as in a Google's TicTacToe doodle
func (_Games) ChooseMarks(g CliGame) (domain.Mark, error) {
	if g.Reader() == nil {
		return "", ErrNilReader()
	}
	read := g.Reader()
	return read(), nil
}

// Local IO

func PrintGame(g CliGame) {
	fmt.Println()

	fmt.Println(g.Player1())
	fmt.Println(g.Player2())

	PrintBoard(g.Board())
}

func PrintBoard(s fmt.Stringer) {
	// Explicit check for the interface
	// var _ fmt.Stringer = brd

	fmt.Println()
	fmt.Println()
	fmt.Println("Press 1 to 9 to mark an empty cell (5 is center), then press ENTER. Board:")
	fmt.Println()

	fmt.Println(s)
	fmt.Println()
}

func PrintWinner(p domain.Player) {
	fmt.Printf("%v won!\n", p)
}

func PrintDraw() {
	fmt.Println("Draw!")
}

// Implicit check for `fmt.Stringer` impl
func Prompt(s fmt.Stringer) { // otherwise `type not defined in this package`
	fmt.Printf("%v, your turn: ", s)
}
