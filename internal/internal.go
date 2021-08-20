package internal

import "context"

type Ctx = context.Context

type Empty = bool

var (
	Yes = true
	No  = false
)

type Id = string

func NewId() Id {
	return "1234"
}
