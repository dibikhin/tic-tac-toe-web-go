package game

import (
	"strings"
	. "tictactoeweb/internal"

	domain "tictactoeweb/internal/domain/game"
)

type (
	CliBoard struct {
		id Id
		domain.Board
		grid
	}
	Row = []string
)

type grid = string

var Gap = domain.Gap

// Factories

func BlankBoard() CliBoard {
	// TODO dup
	row := Row{Gap, Gap, Gap}
	return CliBoard{
		id: Gap,
		grid: strings.Join(Row{
			strings.Join(row, " "),
			strings.Join(row, " "),
			strings.Join(row, " "),
		}, "\n"),
	}
}

func DeadBoard() CliBoard {
	// TODO dup
	row := Row{domain.X_x, domain.X_x, domain.X_x}
	return CliBoard{
		id: domain.X_x,
		grid: strings.Join(Row{
			strings.Join(row, " "),
			strings.Join(row, " "),
			strings.Join(row, " "),
		}, "\n"),
	}
}

func NewBoard(boardId Id, gs ...grid) CliBoard {
	if len(gs) == 1 {
		return CliBoard{
			boardId, domain.NewBoard(), gs[0],
		}
	}
	return CliBoard{
		NewId(), domain.NewBoard(), BlankBoard().grid,
	}
}

// Props

// Checks

func (b CliBoard) IsEmpty() Empty {
	return b.grid == ""
}
