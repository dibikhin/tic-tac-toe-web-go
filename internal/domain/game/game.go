package game

import (
	"fmt"
	irn "tictactoeweb/internal"
)

// Aliasing preventing circular deps
type Reader = func() string

type MakeBoard = func(b ...Board) Board

// Game

type Game struct {
	id irn.Id

	board   Board
	player1 Player
	player2 Player

	reader Reader
}

func NewGame(makeBoard MakeBoard, bs ...Board) Game {
	if len(bs) == 1 {
		return Game{board: bs[0]}
	}
	return Game{
		id: irn.NewId(), board: makeBoard(),
	}
}

// Properties

func (g Game) Id() irn.Id {
	return g.id
}

func (g Game) Board() Board {
	return g.board
}

func (g Game) Player1() Player {
	return g.player1
}

func (g Game) Player2() Player {
	return g.player2
}

// Reader

func (g Game) Reader() Reader {
	return g.reader
}

func (g Game) SetReader(r Reader, def Game) (Game, error) {
	if r == nil {
		return def, irn.ErrNilReader()
	}
	g.reader = r
	return g, nil
}

// IO

func (g Game) Print() {
	fmt.Println()

	fmt.Println(g.player1)
	fmt.Println(g.player2)

	g.board.print()
}

func printWinner(p Player) {
	fmt.Printf("%v won!\n", p)
}

func printDraw() {
	fmt.Println("Draw!")
}
