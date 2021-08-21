package server

import . "tictactoeweb/internal"
import "tictactoeweb/internal/domain"

type Player domain.Player

// Checks

func (p Player) IsEmpty() Empty {
	return p == Player{}
}
