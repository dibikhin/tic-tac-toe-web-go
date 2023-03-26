package gameserver

import (
	"strconv"
	"tictactoe/pkg/api"
	"time"
)

type game struct {
	status api.GameStatus

	id string

	player1   player
	player2   player
	playerWon player
	players   map[name]mark

	board board
}

type player struct {
	mark mark
	name name
}

type mark = string
type name = string

func MakeGame(nam name) game {
	return game{
		status:    api.GameStatus_WAITING_P2_JOIN,
		id:        strconv.Itoa(time.Now().Nanosecond()),
		player1:   player{mark: "X", name: nam},
		player2:   player{},
		playerWon: player{},
		players:   map[name]mark{nam: "X"},
		board:     blankBoard(),
	}
}

func (g game) isEnded() bool {
	return g.status == api.GameStatus_DRAW ||
		g.status == api.GameStatus_WON
}
