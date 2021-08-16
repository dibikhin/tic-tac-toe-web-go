package client

import (
	"errors"

	"tictactoeweb/api"
	"tictactoeweb/internal/domain"
	. "tictactoeweb/internal/domain/game"
)

type again = bool

// Constants, Public
// ErrCouldNotStart arises when `Loop()` is run without running `Setup()` first.
func ErrCouldNotStart() error {
	return errors.New("Game: couldn't start the Game loop, set up the Game first")
}

// Public

// Game Loop()

// Loop prompts players to take turns.
func Loop(g Game) (Game, again, error) {
	if !g.IsReady() {
		return domain.Games.MakeDead(), false, ErrCouldNotStart()
	}
	gam, more := turn(g.Player1(), g)
	if !more {
		return gam, false, nil
	}
	gam, more = turn(gam.Player2(), gam)
	return gam, more, nil
}

// Private

func turn(plr Player, gam Game) (Game, bool) {
	t := takeTurn(plr, gam)
	brd := domain.Boards.Turn(gam.Board(), t)
	printOutcome(brd || remote_game)
	return remote_game
}

func takeTurn(plr Player, gam Game) Turn {
	domain.Prompt(plr)
	for {
		turn := readKey(gam, plr)
		return readTurn(turn, gam, plr)
	}
}

func printOutcome(gam Game) {
	domain.PrintBoard(gam.Board())

	switch o := gam.Outcome(); o {
	case api.Outcome_DRAW:
		domain.PrintDraw()

	case api.Outcome_WON:
		domain.PrintWinner(gam.Winner())
	}

	// Party:Server

	// if brd.IsWinner(plr.Mark()) {
	// 	domain.PrintWinner(plr)
	// 	return domain.Games.SetBoard(gam, brd), false
	// }
	// if !brd.HasEmpty() {
	// 	domain.PrintDraw()
	// 	return domain.Games.SetBoard(gam, brd), false
	// }
	// return domain.Games.SetBoard(gam, brd), true
}

func readKey(g Game, plr Player) Key {
	for {
		read := g.Reader()
		turn := Key(read())

		// Party: Server ?
		if !turn.IsKey() {
			domain.PrintBoard(g.Board())
			domain.Prompt(plr)

			continue
		}
		return turn
	}
}

func readTurn(turn Key, g Game, plr Player) Turn {
	for {
		cel := turn.ToCell()
		if g.Board().IsFilled(cel) {
			domain.PrintBoard(g.Board())
			domain.Prompt(plr)

			continue
		}
		return NewTurn(plr.Mark(), cel)
	}
}
