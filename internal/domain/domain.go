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

// Party: Server
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

func (_Boards) Turn(b Board, t Turn) Board {
	// WARN: possible out of range
	// b[c.row][c.col] = m
	// TODO: send to server

	return Board{}
}

// Commands: Local + IO

// ChooseMarks chooses players' marks as in a Google's TicTacToe doodle
func (_Games) ChooseMarks(g Game) (Player, Player /*Game? TODO:*/, error) {
	fmt.Print("Press 'x' or 'o' to choose mark for Player 1: ")

	if g.Reader() == nil {
		return DeadPlayer(), DeadPlayer(), irn.ErrNilReader()
	}
	read := g.Reader()
	m := read()
	p1, p2 := Games.ArrangePlayers(m)
	return p1, p2, nil
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
