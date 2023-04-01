package gameclient

import (
	"fmt"
	"tictactoe/pkg/api"
)

type Mark string
type Name string

type Player struct {
	mark Mark
	name Name
}

func NewPlayer(r *api.Player) Player {
	return Player{
		mark: Mark(r.Mark),
		name: Name(r.Name),
	}
}

func (p *Player) String() string {
	return fmt.Sprintf("name: %v, mark: %v", p.name, p.mark)
}
