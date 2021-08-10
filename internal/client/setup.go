package client

import (
	"fmt"
	"tictactoeweb/internal/domain"
	game "tictactoeweb/internal/domain/game"
)

// Setup initializes the game and helps players to choose marks.
// The param is a strategy for user input to be stubbed.
// One can pass nothing, the default reader is used in the case.
// Example:
// ctx, err := Setup()
// OR
// ctx, err := Setup(DefaultReader)
// OR
// ctx, err := Setup(yourReaderFunc)
func Setup(rs ...Reader) (game.Game, error) {
	alt, err := extractReader(rs)
	if err != nil {
		return domain.Games.Dead(), err
	}
	gam, err := makeGame(DefaultReader, alt)
	if err != nil {
		return domain.Games.Dead(), err
	}
	printLogo(game.Logo())

	defer gam.Print()

	p1, p2, err := domain.Games.ChooseMarks(gam)
	if err != nil {
		return domain.Games.Dead(), err
	}
	return domain.Games.SetPlayers(gam, p1, p2), nil
}

// IO

func printLogo(s fmt.Stringer) {
	fmt.Println()
	fmt.Println(s)
	fmt.Println()

	fmt.Println("(Use `ctrl+c` to exit)")
	fmt.Println()
}

// Private

// Factory, Pure
func makeGame(def, alt Reader) (game.Game, error) {
	g := domain.Games.NewGame()
	switch {
	case alt != nil:
		return domain.Games.SetReader(g, alt)
	default:
		return g.SetReader(def, domain.Games.Dead())
	}
}
