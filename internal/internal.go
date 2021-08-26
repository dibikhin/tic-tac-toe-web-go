package internal

import "context"

type Empty = bool

const (
	Yes = true
	No  = false
)

type Id = string

func NewId() Id {
	return "1234"
}

type Ctx = context.Context
