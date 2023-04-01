package gameserver

import "strconv"

type Key string

type Cell struct {
	row, col int
}

func (k Key) isKey() bool {
	n, err := strconv.Atoi(string(k))
	if err != nil {
		return false
	}
	return n >= 1 && n <= 9
}

func (k Key) toCell() Cell {
	return coords()[k]
}

func coords() map[Key]Cell {
	return map[Key]Cell{
		"1": {0, 0}, "2": {0, 1}, "3": {0, 2},
		"4": {1, 0}, "5": {1, 1}, "6": {1, 2},
		"7": {2, 0}, "8": {2, 1}, "9": {2, 2},
	}
}
