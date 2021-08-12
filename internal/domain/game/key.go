package game

import "strconv"

// Key

type (
	Key    string
	coords map[Key]Cell
)

// Cell

type Cell struct {
	row, col int
}

// Public

// Props

func (c Cell) Row() int {
	return c.row
}

func (c Cell) Col() int {
	return c.col
}

// Other

func (k Key) IsKey() bool {
	n, err := strconv.Atoi(string(k))
	if err != nil {
		return false
	}
	return n >= 1 && n <= 9
}

func (k Key) ToCell() Cell {
	return _coords()[k] // TODO: detect and propagate errors?
}

// Constants, Private

func _coords() coords {
	return coords{
		"1": {0, 0}, "2": {0, 1}, "3": {0, 2},
		"4": {1, 0}, "5": {1, 1}, "6": {1, 2},
		"7": {2, 0}, "8": {2, 1}, "9": {2, 2},
	}
}
