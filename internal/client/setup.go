package client

import (
	. "tictactoeweb/internal"
	"tictactoeweb/internal/domain/game"

	. "tictactoeweb/internal/client/game"
)

var MakeDeadGame = Domain.Games.MakeDead

// Constants
func Logo() game.Board {
	return game.NewBoard(
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
func SetupMarks(rs ...Reader) (CliGame, error) {
	Domain.Greet(Logo())
	Domain.PromptMark()

	gm := Domain.Games.Make()
	game, err := chooseMarks(gm)

	defer Domain.PrintGame(game)
	return game, err
}

// Private

func chooseMarks(game CliGame) (CliGame, error) {
	mark, err := Domain.Games.ReadMark(game)
	if err != nil {
		return game, err
	}
	gm, err := Domain.Games.ArrangePlayers(mark)
	return gm, err
}
