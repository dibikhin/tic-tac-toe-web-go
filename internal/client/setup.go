package client

import (
	"fmt"

	. "tictactoeweb/internal"

	. "tictactoeweb/internal/client/game"
)

var MakeDeadGame = Domain.Games.MakeDead

// Constants
func Logo() CliBoard {
	return NewCliBoard(
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
func Setup(rs ...Reader) (CliGame, error) {
	alt, err := ExtractReader(rs...)
	if err != nil {
		return MakeDeadGame(), err
	}
	game, err := setupReader(DefaultReader, alt)
	if err != nil {
		return game, err
	}
	printLogo(Logo())
	Domain.PromptMark()

	mark, err := Domain.Games.ChooseMarks(game)
	if err != nil {
		return game, err
	}
	gm, err := Domain.Games.ArrangePlayers(mark)
	if err != nil {
		return gm, err
	}
	defer Domain.PrintGame(gm)
	return gm, nil
}

// Private

// Factory
func setupReader(defualt, alt Reader) (CliGame, error) {
	gm := Domain.Games.Make()
	switch {
	case alt != nil:
		return gm.SetReader(alt, MakeDeadGame())
	default:
		return gm.SetReader(defualt, MakeDeadGame())
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
