package game

import (
	"fmt"

	. "tictactoeweb/internal"
)

type Player struct {
	mark Mark
	num  int32 // 1 or 2; -1 is a dead Player
}

// Factorys

func NoPlayer() Player {
	return NewPlayer(Gap, -1)
}

func DeadPlayer() Player {
	return NewPlayer(X_x, -1)
}

func NewPlayer(m Mark, n int32) Player {
	return Player{m, n}
}

// Public

func (p Player) String() string {
	return fmt.Sprintf(`Player %v ("%v")`, p.num, p.mark)
}

// Props

func (p Player) Mark() Mark {
	return p.mark
}

func (p Player) Num() int32 {
	return p.num
}

// Checks

func (p Player) IsEmpty() Empty {
	return p == Player{} || p == NoPlayer()
}
