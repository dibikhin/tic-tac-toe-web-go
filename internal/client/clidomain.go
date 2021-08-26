package client

import (
	"fmt"

	. "tictactoeweb/internal"
	. "tictactoeweb/internal/client/game"
	. "tictactoeweb/internal/domain/game"
)

type (
	_Domain struct {
		Games // see the `Games` var below
		Boards
	}
	Games  struct{}
	Boards struct{}

	reader = Reader
)

// Public

var Domain = _Domain{} // to call like `Domain.Games.ArrangePlayers(m)`

// Factorys

func (Games) Make() CliGame {
	return NewCliGame(New Id())
}

func (Games) MakeDead() CliGame {
	return NewCliGame{
		Game: NewGame(X_x),
	}
}

// IO

// Commands: Local

// ChooseMarks chooses players' marks as in a Google's TicTacToe doodle
func (Games) ChooseMarks(g CliGame) (Mark, error) {
	if g.Reader() == nil {
		return "", ErrNilReader()
	}
	read := g.Reader()
	return read(), nil
}

// Commands: Remote

func (Games) ArrangePlayers(m Mark) (CliGame, error) {
	return CliGame{}, nil
}

func (Boards) Turn(boardId Id, trn Turn) (CliGame, error) {
	return CliGame{}, nil
}

// Querys: Remote
func (Boards) IsFilled(boardId Id, key CliKey) (bool, error) {
	return Yes, nil
}

// Local IO

func (_Domain) promptMark() {
	fmt.Print("Press 'x' or 'o' to choose mark for Player 1: ")
}

func (_Domain) PrintGame(g CliGame) {
	fmt.Println()

	fmt.Println(g.Player1())
	fmt.Println(g.Player2())

	Domain.PrintBoard(g.Board())
}

func (_Domain) PrintBoard(str fmt.Stringer) {
	// Explicit check for the interface
	// var _ fmt.Stringer = brd

	fmt.Println()
	fmt.Println()
	fmt.Println("Press 1 to 9 to mark an empty cell (5 is center), then press ENTER. Board:")
	fmt.Println()

	fmt.Println(str)
	fmt.Println()
}

func (_Domain) PrintWinner(s fmt.Stringer) {
	fmt.Printf("%v won!\n", s)
}

func (_Domain) PrintDraw() {
	fmt.Println("Draw!")
}

// Implicit check for `fmt.Stringer` impl
func (_Domain) Prompt(s fmt.Stringer) { // otherwise `type not defined in this package`
	fmt.Printf("%v, your turn: ", s)
}
