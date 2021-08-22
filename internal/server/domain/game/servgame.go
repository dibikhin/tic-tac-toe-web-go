package game

import (
	domain "tictactoeweb/internal/domain/game"
)

type ServGame struct {
	domain.Game

	board ServBoard
}

// Checks

func (g ServGame) IsReady() bool {
	return !g.Player1().IsEmpty() &&
		!g.Player2().IsEmpty() &&
		!g.Board().IsEmpty()
}
