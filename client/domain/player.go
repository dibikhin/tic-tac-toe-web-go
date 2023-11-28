package domain

import (
	"fmt"
	"tictactoe/api"
)

const StartGame Command = "p"

type (
	Mark       string
	PlayerName string
	Turn       string
	Command    string

	Player struct {
		Mark Mark
		Name PlayerName
	}
)

func ToPlayer(r *api.Player) Player {
	return Player{
		Mark: Mark(r.Mark),
		Name: PlayerName(r.Name),
	}
}

func (p *Player) String() string {
	if p == nil || p.Name == "" || p.Mark == "" {
		return ""
	}
	return fmt.Sprintf("name: %v, mark: %v", p.Name, p.Mark)
}
