package internal

import "context"

// Anti-pattern :)

type Ctx = context.Context

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
