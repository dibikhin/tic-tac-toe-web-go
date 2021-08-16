package client

import (
	"fmt"

	irn "tictactoeweb/internal"
	"tictactoeweb/internal/domain"
	"tictactoeweb/internal/domain/game"
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
func Setup(rs ...irn.Reader) (game.Game, error) {
	alt, err := irn.ExtractReader(rs)
	if err != nil {
		return domain.Games.MakeDead(), err
	}
	gam, err := setupReader(irn.DefaultReader, alt)
	if err != nil {
		return domain.Games.MakeDead(), err
	}
	printLogo(domain.Logo())

	defer domain.PrintGame(gam)
	p1, p2, err := domain.Games.ChooseMarks(gam)
	if err != nil {
		return domain.Games.MakeDead(), err
	}
	return domain.Games.SetPlayers(gam, p1, p2), nil
}

// Private

// Factory, Pure
func setupReader(def, alt irn.Reader) (game.Game, error) {
	dead := domain.Games.MakeDead()
	gam := domain.Games.Make()
	switch {
	case alt != nil:
		return gam.SetReader(alt, dead)
	default:
		return gam.SetReader(def, dead)
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
