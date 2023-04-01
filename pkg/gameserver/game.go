package gameserver

import (
	"strconv"
	"tictactoe/pkg/api"
	"time"
)

type Game struct {
	status api.GameStatus

	id ID

	player1   Player
	player2   Player
	playerWon Player
	players   map[Name]Mark

	board Board
}

type ID string

type Player struct {
	mark Mark
	name Name
}

type Mark string
type Name string

func MakeGame(name Name) Game {
	return Game{
		status:    api.GameStatus_WAITING_P2_JOIN,
		id:        ID(strconv.Itoa(time.Now().Nanosecond())),
		player1:   Player{mark: "X", name: name},
		player2:   Player{},
		playerWon: Player{},
		players:   map[Name]Mark{name: "X"},
		board:     blankBoard(),
	}
}

func (g Game) isEnded() bool {
	return g.status == api.GameStatus_DRAW ||
		g.status == api.GameStatus_WON
}
