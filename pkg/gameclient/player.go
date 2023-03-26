package gameclient

import "fmt"

type mark = string
type name = string

type player struct {
	mark mark
	name name
}

func (p *player) String() string {
	return fmt.Sprintf("name: %v, mark: %v", p.name, p.mark)
}
