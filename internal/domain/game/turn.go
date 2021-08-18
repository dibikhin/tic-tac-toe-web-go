package game

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
	Len = int  // 1..3
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

// "3" IsIn {"1", "3", "5"}
func (k Key) IsIn(kk []Key) bool {
	for _, v := range kk {
		if v == k {
			return true
		}
	}
	return false
}
