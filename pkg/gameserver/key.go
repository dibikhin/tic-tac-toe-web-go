package gameserver

import "strconv"

type Key string

type Cell struct {
	row, col int
}

func keysCells() map[Key]Cell {
	return map[Key]Cell{
		"1": {0, 0}, "2": {0, 1}, "3": {0, 2},
		"4": {1, 0}, "5": {1, 1}, "6": {1, 2},
		"7": {2, 0}, "8": {2, 1}, "9": {2, 2},
	}
}

func (k Key) isKey() bool {
	n, err := strconv.Atoi(string(k))
	if err != nil {
		return false
	}
	return n >= 1 && n <= 9
}

func (k Key) toCell() Cell {
	return keysCells()[k]
}

func (c Cell) isInRange(b Board) bool {
	return c.row < len(b) && c.col < len(b[c.row])
}
