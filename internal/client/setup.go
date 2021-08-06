package client

import (
	"fmt"
	. "tictactoe/internal"
	"tictactoeweb/internal/domain"
	"tictactoeweb/internal/domain/game"
)

// Constants, Private


// Setup initializes the game and helps players to choose marks.
// The param is a strategy for user input to be stubbed.
// One can pass nothing, the default reader is used in the case.
// Example:
// ctx, err := Setup()
// OR
// ctx, err := Setup(DefaultReader)
// OR
// ctx, err := Setup(yourReaderFunc)
func Setup(rs ...reader) (game, error) {
	alt, err := extractReader(rs)
	if err != nil {
		return DeadGame(), err
	}
	gam, err := makeGame(DefaultReader, alt)
	if err != nil {
		return DeadGame(), err
	}
	printLogo(domain.Board.Logo())

	defer gam.Print()
	p1, p2, err := gam.ChooseMarks()
	if err != nil {
		return DeadGame(), err
	}
	return SetPlayers(gam, p1, p2), nil
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
func makeGame(def, alt reader) (Game, error) {
	g := domain.Game.NewGame()
	switch {
	case alt != nil:
		return domain.Game.SetReader(g, alt)
	default:
		return SetReader(g, def)
	}
}
