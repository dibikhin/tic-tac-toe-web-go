package game

import (
	"fmt"
	"tictactoe/api"
)

type Mark string
type Name string

type Player struct {
	Mark Mark
	Name Name
}

func ToPlayer(r *api.Player) Player {
	return Player{
		Mark: Mark(r.Mark),
		Name: Name(r.Name),
	}
}

func (p *Player) String() string {
	return fmt.Sprintf("name: %v, mark: %v", p.Name, p.Mark)
}
