package internal

import (
	"fmt"
	"strings"
)

// User input strategy for stubbing in tests.
//
// NOTE: An interface is more idiomatic in this case. BUT it's overkill to define
// a type with constructor, an interface and its fake implementation in tests vs. this
// func, its impl and its fake impl in tests.
type reader func() string

// Game

type game struct {
	logo board

	board   board
	player1 player
	player2 player

	read reader
}

// Constants

func _deadGame() game {
	return game{board: _deadBoard()}
}

// Public

// Pure
func (g game) Board() board {
	return g.board
}

// Private

// Pure
func newGame() game {
	return game{
		logo:  _logo(),
		board: _blankBoard(),

		// others are omitted for flexibility
	}
}

// Setters

func setBoard(g game, b board) game {
	g.board = b
	return g
}

func setPlayers(g game, p1, p2 player) game {
	g.player1 = p1
	g.player2 = p2
	return g
}

func setReader(g game, read reader) game {
	g.read = read // WARN possible nil
	return g
}

// Pure
func (g game) isReady() bool {
	return g.read != nil &&
		!g.player1.isEmpty() &&
		!g.player2.isEmpty() &&
		!g.board.isEmpty()
}

// Setup() IO

func printLogo(s fmt.Stringer) {
	fmt.Println()
	fmt.Println(s)
	fmt.Println()

	fmt.Println("(Use `ctrl+c` to exit)")
	fmt.Println()
}

func (g game) chooseMarks() (player, player) {
	fmt.Print("Press 'x' or 'o' to choose mark for Player 1: ")

	// WARN: possible nil, ignored
	mark1 := g.read()
	p1, p2 := arrangePlayers(mark1)
	return p1, p2
}

func (g game) print() {
	fmt.Println()

	fmt.Println(g.player1)
	fmt.Println(g.player2)

	g.board.print()
}

// Other

// Pure
func arrangePlayers(m mark) (player, player) {
	if strings.ToLower(m) == "x" {
		return player{"X", 1}, player{"O", 2}
	}
	return player{"O", 1}, player{"X", 2}
}
