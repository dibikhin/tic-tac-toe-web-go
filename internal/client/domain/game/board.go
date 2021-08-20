package game

import (
	"strings"

	. "tictactoeweb/internal"
	. "tictactoeweb/internal/domain"
)

type (
	Board struct {
		id Id
		Board
		grid
	}
	Row = [Size]string
)

type grid = string

// Factories

func BlankBoard() Board {
	// TODO dup
	row := Row{Gap, Gap, Gap}
	return Board{
		id: Gap,
		grid: strings.Join(Row{
			strings.Join(row, " "),
			strings.Join(row, " "),
			strings.Join(row, " "),
		}, "\n"),
	}
}

func DeadBoard() Board {
	// TODO dup
	row := Row{X_x, X_x, X_x}
	return Board{
		id: X_x,
		grid: strings.Join(Row{
			strings.Join(row, " "),
			strings.Join(row, " "),
			strings.Join(row, " "),
		}, "\n"),
	}
}

func NewBoard(boardId Id, gs ...grid) Board {
	if len(gs) == 1 {
		return Board{
			boardId,
			NewBoard(),
			gs[0],
		}
	}
	return Board{
		NewId(), NewBoard(), BlankBoard().grid,
	}
}

// Props

func (b Board) Id() Id {
	return b.id
}

// Checks

func (b Board) IsEmpty() Empty {
	return b.grid == ""
}
