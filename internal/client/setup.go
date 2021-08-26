package client

import (
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
	Domain.Greet(Logo())

	game, err := prepareGame(rs...)
	if err != nil {
		return game, err
	}
	Domain.PromptMark()

	gm, err := chooseMarks(game)
	defer Domain.PrintGame(gm)
	return gm, err
}

// Private

func prepareGame(rs ...Reader) (CliGame, error) {
	alt, err := ExtractReader(rs...)
	if err != nil {
		return MakeDeadGame(), err
	}
	g := Domain.Games.Make()
	game, err := setupReader(g, DefaultReader, alt)
	return game, err
}

func chooseMarks(game CliGame) (CliGame, error) {
	mark, err := Domain.Games.ReadMark(game)
	if err != nil {
		return game, err
	}
	gm, err := Domain.Games.ArrangePlayers(mark)
	return gm, err
}

// Factory
func setupReader(gm CliGame, defualt, alt Reader) (CliGame, error) {
	switch {
	case alt != nil:
		return gm.SetReader(alt, MakeDeadGame())
	default:
		return gm.SetReader(defualt, MakeDeadGame())
	}
}
