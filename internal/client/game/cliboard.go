package game

import (
	"fmt"

	. "tictactoeweb/internal"

	. "tictactoeweb/internal/domain/game"
)

type (
	CliBoard struct {
		id Id
		Board
		grid
	}
)

type grid string

// Factories

func BlankCliBoard() CliBoard {
	return NewCliBoard(
		Gap, grid(composeGrid(Gap)),
	)
}

func DeadCliBoard() CliBoard {
	return NewCliBoard(
		X_x, grid(composeGrid(X_x)),
	)
}

func NewCliBoard(id Id, gs ...grid) CliBoard {
	if len(gs) == 1 {
		return CliBoard{
			id, NewBoard(NewId()), gs[0],
		}
	}
	return CliBoard{
		id: NewId(), Board: NewBoard(NewId()),
	}
}

// Props

// Checks

func (b CliBoard) IsEmpty() Empty {
	return b.Id() == "" || b.grid.IsEmpty() || b == BlankCliBoard()
}

func (g grid) IsEmpty() Empty {
	return g == ""
}

// Private

func composeGrid(s string) string {
	return fmt.Sprintf(
		`%v %v %v
%v %v %v
%v %v %v`,
		s, s, s,
		s, s, s,
		s, s, s,
	)
}
