package internal

import "fmt"

type player struct {
	mark mark
	num  int // 1 or 2; -1 is a dead player
}

func (p player) String() string {
	return fmt.Sprintf(`Player %v ("%v")`, p.num, p.mark)
}

func _deadPlayer() player {
	return player{"X_x", -1}
}

func (p player) isEmpty() bool {
	return p == player{}
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
