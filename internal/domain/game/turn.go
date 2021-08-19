package game

import . "tictactoeweb/internal"

// Public
type (
	Turn struct {
		key  Key
		mark Mark
	}
	Key string // "1".."9"
)

// Public

func NewTurn(m Mark, k Key) Turn {
	return Turn{k, m}
}

func NoTurn() Turn {
	return Turn{"-1", "No"}
}

// "3" IsIn {"1", "3", "5"}
func (k Key) IsIn(kk []Key) bool {
	for _, v := range kk {
		if v == k {
			return Yes
		}
	}
	return No
}
