package client

import "fmt"

type mark = string

type player struct {
	mark mark
	name string
}

func (p *player) String() string {
	return fmt.Sprintf("name: %v, mark: %v", p.name, p.mark)
}
