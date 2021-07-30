package internal

import (
	"errors"
	"fmt"
	"strings"
)

// Constants, Private
// ErrNilReader() arises when `Setup()` is run with nil reader.
func ErrNilReader() error {
	return errors.New("game: the reader is nil, use a non-nil reader or nothing for the default one while setting up")
}

type reader = func() string

// Game

type Game struct {
	board   Board
	player1 player
	player2 player

	read reader
}

// Constants

func DeadGame() Game {
	return Game{board: _deadBoard()}
}

// Public

// Pure
func (g Game) Board() Board {
	return g.board
}

// Private

// Pure
func NewGame() Game {
	return Game{
		board: _blankBoard(),

		// the rest fields are omitted for flexibility
	}
}

// Setters

func SetPlayers(g Game, p1, p2 player) Game {
	g.player1 = p1
	g.player2 = p2
	return g
}

func SetReader(g Game, r reader) (Game, error) {
	if r == nil {
		return DeadGame(), ErrNilReader()
	}
	g.read = r
	return g, nil
}

func setBoard(g Game, b Board) Game {
	g.board = b
	return g
}

// Pure
func (g Game) isReady() bool {
	return g.read != nil &&
		!g.player1.isEmpty() &&
		!g.player2.isEmpty() &&
		!g.board.isEmpty()
}

// IO

func (g Game) ChooseMarks() (player, player, error) {
	fmt.Print("Press 'x' or 'o' to choose mark for Player 1: ")

	if g.read == nil {
		return _deadPlayer(), _deadPlayer(), ErrNilReader()
	}
	mark1 := g.read()
	p1, p2 := arrangePlayers(mark1)
	return p1, p2, nil
}

func (g Game) Print() {
	fmt.Println()

	fmt.Println(g.player1)
	fmt.Println(g.player2)

	g.board.print()
}

func printWinner(p player) {
	fmt.Printf("%v won!\n", p)
}

func printDraw() {
	fmt.Println("Draw!")
}

// Other

// Pure
func arrangePlayers(m mark) (player, player) {
	if strings.ToLower(m) == "x" {
		return player{"X", 1}, player{"O", 2}
	}
	return player{"O", 1}, player{"X", 2}
}
