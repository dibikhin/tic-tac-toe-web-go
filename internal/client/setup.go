package client

import (
	"fmt"

	irn "tictactoeweb/internal"
	"tictactoeweb/internal/domain"
	. "tictactoeweb/internal/domain/game"
)

// Constants
func Logo() Board {
	return NewBoard(
		"logo",
		[Size][Size]Mark{
			{X, " ", X},
			{O, X, O},
			{X, " ", O}},
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
func Setup(rs ...irn.Reader) (Game, error) {
	dead := domain.Games.MakeDead
	alt, err := irn.ExtractReader(rs)
	if err != nil {
		return dead(), err
	}
	gam, err := setupReader(irn.DefaultReader, alt)
	if err != nil {
		return dead(), err
	}
	defer domain.PrintGame(gam)
	printLogo(Logo())
	promptMark()

	mrk, err := domain.Games.ChooseMarks(gam)
	if err != nil {
		return dead(), err
	}
	g, err := domain.Games.ArrangePlayers(mrk)
	if err != nil {
		return dead(), err
	}
	return g, nil
}

// Private

// Factory, Pure
func setupReader(def, alt irn.Reader) (Game, error) {
	dead := domain.Games.MakeDead
	gam := domain.Games.Make()
	switch {
	case alt != nil:
		return gam.SetReader(alt, dead())
	default:
		return gam.SetReader(def, dead())
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
