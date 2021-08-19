package game

import (
	"strings"

	. "tictactoeweb/internal"
)

type (
	Board struct {
		id Id
		grid
	}
	Row  = [Size]string
	Mark = string // "X" or "O" (or "No")
)

type grid = string

// Constants, Public

const (
	Size = 3
	Gap  = __
	X    = "X"
	O    = "O"
	X_x  = "X_x"
)

const __ = "-" // Private

// Factories

func BlankBoard() Board {
	row := Row{__, __, __}
	rows := Row{
		strings.Join(row[:], " "),
		strings.Join(row[:], " "),
		strings.Join(row[:], " "),
	}
	return Board{
		id:   __,
		grid: strings.Join(rows[:], "\n"),
	}
}

func Dead() Board {
	row := Row{X_x, X_x, X_x}
	rows := Row{
		strings.Join(row[:], " "),
		strings.Join(row[:], " "),
		strings.Join(row[:], " "),
	}
	return Board{
		id:   X_x,
		grid: strings.Join(rows[:], "\n"),
	}
}

// Public

func NewBoard(boardId Id, gs ...grid) Board {
	if len(gs) == 1 {
		return Board{
			boardId, gs[0],
		}
	}
	return Board{
		NewId(), BlankBoard().grid,
	}
}

// Props

func (b Board) Id() Id {
	return b.id
}

func (b Board) IsEmpty() Empty {
	return b.grid == ""
}

func (b Board) String() string {
	return b.grid
}
