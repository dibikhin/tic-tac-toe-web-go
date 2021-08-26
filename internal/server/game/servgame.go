package game

import (
	. "tictactoeweb/internal/domain/game"
)

type ServGame struct {
	Game

	board ServBoard
}
