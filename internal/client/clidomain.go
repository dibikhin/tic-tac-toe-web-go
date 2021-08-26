package client

import (
	"fmt"

	. "tictactoeweb/internal"
	. "tictactoeweb/internal/client/game"
	. "tictactoeweb/internal/domain/game"
)

type (
	_Domain struct {
		Games  _Games // see the `Games` var below
		Boards _Boards
	}
	_Games  struct{}
	_Boards struct{}
)

// Public

// Globals

var Domain = _Domain{} // to call like `Domain.Games.ArrangePlayers(m)`

// Factorys

func (_Games) Make() CliGame {
	return NewCliGame(NewId())
}

func (_Games) MakeDead() CliGame {
	return NewCliGame(X_x)
}

// IO

// Commands: Local

// ReadMark chooses players' marks as in a Google's TicTacToe doodle
func (_Games) ReadMark(g CliGame) (Mark, error) {
	if App.Reader() == nil {
		return DeadMark(), ErrNilReader()
	}
	read := App.Reader()
	return read(), nil
}

// Commands: Remote

func (_Games) ArrangePlayers(m Mark) (CliGame, error) {
	return CliGame{}, nil
}

func (_Boards) Turn(boardId Id, trn Turn) (CliGame, error) {
	return CliGame{}, nil
}

// Querys: Remote
func (_Boards) IsFilled(boardId Id, key CliKey) (bool, error) {
	return Yes, nil
}

// Local IO

func (_Domain) Greet(str fmt.Stringer) {
	fmt.Println("Hi!")
	fmt.Println()
	fmt.Println("This is 3x3 Tic-tac-toe for 2 friends :)")
	fmt.Println()
	fmt.Println(str)
	fmt.Println()

	fmt.Println("(Use `ctrl+c` to exit)")
	fmt.Println()
}

func (_Domain) PromptMark() {
	fmt.Print("Press 'x' or 'o' to choose mark for Player 1, then press ENTER: ")
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
