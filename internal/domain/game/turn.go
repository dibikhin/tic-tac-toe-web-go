package game

// Public
type (
	Turn struct {
		key  Key
		mark Mark
	}
	Key string // "1".."9" | "No"
)

// Public

func NewTurn(m Mark, k Key) Turn {
	return Turn{k, m}
}

func NoTurn() Turn {
	return NewTurn("-1", "No")
}
