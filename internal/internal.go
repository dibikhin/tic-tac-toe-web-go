package internal

// Anti-pattern :)

// Bools

type Empty = bool

const (
	Yes = true
	No  = false
)

// Id

type Id = string

func NewId() Id {
	return "1234"
}
