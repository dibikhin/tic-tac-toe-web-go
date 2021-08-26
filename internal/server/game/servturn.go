package game

import (
	"strconv"
	. "tictactoeweb/internal/domain/game"
)

// Public

type (
	ServKey Key
	Cell    struct {
		row, col Len
	}
	Len = int // 1..3
)

// Private

type (
	coords map[ServKey]Cell
)

// Other

func (k ServKey) IsKey() bool {
	n, err := strconv.Atoi(string(k))
	if err != nil {
		return false
	}
	return n >= 1 && n <= 9
}

// Constants, Private

func _coords() coords {
	return coords{
		"1": {0, 0}, "2": {0, 1}, "3": {0, 2},
		"4": {1, 0}, "5": {1, 1}, "6": {1, 2},
		"7": {2, 0}, "8": {2, 1}, "9": {2, 2},
	}
}

func (k ServKey) ToCell() Cell {
	return _coords()[k] // TODO: detect and propagate errors?
}

// Props

func (c Cell) Row() int {
	return c.row
}

func (c Cell) Col() int {
	return c.col
}
