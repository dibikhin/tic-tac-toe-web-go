package domain

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
