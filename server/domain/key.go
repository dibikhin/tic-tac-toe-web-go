package domain

type Cell struct {
	row, col int
}

func (c Cell) isInRange(b Board) bool {
	return c.row < len(b) && c.col < len(b[c.row])
}

func keyToCell() map[Key]Cell {
	return map[Key]Cell{
		"1": {0, 0}, "2": {0, 1}, "3": {0, 2},
		"4": {1, 0}, "5": {1, 1}, "6": {1, 2},
		"7": {2, 0}, "8": {2, 1}, "9": {2, 2},
	}
}
