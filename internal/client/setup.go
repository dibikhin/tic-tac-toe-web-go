package client

import (
	"fmt"

	. "tictactoeweb/internal"
	"tictactoeweb/internal/domain"
	. "tictactoeweb/internal/domain/game"
)

var deadGame = domain.Games.MakeDead

// Constants
func Logo() Board {
	return NewBoard(
		"logo",
		`O   X
	     O X X
	     X   O`,
	)
}

// Setup initializes the game and helps players to choose marks.
// The param is a strategy for user input to be stubbed.
// One can pass nothing, the default reader is used in the case.
// Example:
// ctx, err := Setup()
// OR
// ctx, err := Setup(DefaultReader)
// OR
// ctx, err := Setup(yourReaderFunc)
func Setup(rs ...Reader) (Game, error) {
	alt, err := ExtractReader(rs...)
	if err != nil {
		return deadGame(), err
	}
	gam, err := setupReader(DefaultReader, alt)
	if err != nil {
		return deadGame(), err
	}
	defer domain.PrintGame(gam)
	printLogo(Logo())
	promptMark()

	mrk, err := domain.Games.ChooseMarks(gam)
	if err != nil {
		return deadGame(), err
	}
	g, err := domain.Games.ArrangePlayers(mrk)
	if err != nil {
		return deadGame(), err
	}
	return g, nil
}

// Private

// Factory
func setupReader(def, alt Reader) (Game, error) {
	gam := domain.Games.Make()
	switch {
	case alt != nil:
		return gam.SetReader(alt, deadGame())
	default:
		return gam.SetReader(def, deadGame())
	}
}

// IO

func printLogo(s fmt.Stringer) {
	fmt.Println()
	fmt.Println(s)
	fmt.Println()

	fmt.Println("(Use `ctrl+c` to exit)")
	fmt.Println()
}

func promptMark() {
	fmt.Print("Press 'x' or 'o' to choose mark for Player 1: ")
}
