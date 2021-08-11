package game

import "fmt"

type Player struct {
	mark Mark
	num  int // 1 or 2; -1 is a dead Player
}

// Public, Pure

func (p Player) String() string {
	return fmt.Sprintf(`Player %v ("%v")`, p.Num(), p.Mark())
}

// Factories

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

// IO

// Implicit check for `fmt.Stringer` impl
func prompt(s fmt.Stringer) { // otherwise `type not defined in this package`
	fmt.Printf("%v, your turn: ", s)
}

func (b Board) print() {
	// Explicit check for the interface
	var _ fmt.Stringer = b

	fmt.Println()
	fmt.Println()
	fmt.Println("Press 1 to 9 to mark an empty cell (5 is center), then press ENTER. Board:")
	fmt.Println()

	fmt.Println(b)
	fmt.Println()
}
