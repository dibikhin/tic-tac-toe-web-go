package game

import "fmt"

type Player struct {
	mark Mark
	num  int // 1 or 2; -1 is a dead Player
}

// Public, Pure

func (p Player) String() string {
	return fmt.Sprintf(`Player %v ("%v")`, p.num, p.mark)
}

// Factorys

func NewPlayer(m Mark, n int) Player {
	return Player{m, n}
}

func DeadPlayer() Player {
	return Player{X_x, -1}
}

// Props

func (p Player) Mark() Mark {
	return p.mark
}

func (p Player) Num() int {
	return p.num
}

// Checks

func (p Player) IsEmpty() bool {
	return p == Player{}
}
