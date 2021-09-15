package client

import (
	. "tictactoeweb/internal"
	"tictactoeweb/internal/domain/game"

	. "tictactoeweb/internal/client/game"
)

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
func SetupMarks(ctx Ctx) (CliGame, error) {
	Domain.Greet(Logo())
	Domain.PromptMark()

	gm := Domain.Games.Make()
	game, err := chooseMarks(ctx, gm)

	defer Domain.PrintGame(game)
	return game, err
}

// Private

func chooseMarks(ctx Ctx, game CliGame) (CliGame, error) {
	mark, err := Domain.Games.ReadMark()
	if err != nil {
		return game, err
	}
	return Domain.Games.ArrangePlayers(ctx, mark)
}
