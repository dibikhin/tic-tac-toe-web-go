package game

import "strconv"

// Public
type (
	Turn struct {
		cel  Cell
		mark Mark
	}
	Cell struct {
		row, col Len
	}
	Key string // "1".."9"
	Len = int    // 1..3
)

// Private

type (
	coords map[Key]Cell
)

// Public

func NewTurn(m Mark, c Cell) Turn {
	return Turn{c, m}
}

func NoTurn() Turn {
	return Turn{Cell{-1, -1}, "No"}
}

func (k Key) ToCell() Cell {
	return _coords()[k] // TODO: detect and propagate errors?
}

// Props

func (c Cell) Row() int {
	return c.row
}

func (c Cell) Col() int {
	return c.col
}

// Other

// Party:Server
func (k Key) IsKey() bool {
	n, err := strconv.Atoi(string(k))
	if err != nil {
		return false
	}
	return n >= 1 && n <= 9
}

// Constants, Private

// Party:Server
func _coords() coords {
	return coords{
		"1": {0, 0}, "2": {0, 1}, "3": {0, 2},
		"4": {1, 0}, "5": {1, 1}, "6": {1, 2},
		"7": {2, 0}, "8": {2, 1}, "9": {2, 2},
	}
}
