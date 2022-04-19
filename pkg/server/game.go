package server

import (
	"strconv"
	"tictactoe/pkg/api"
	"time"
)

type Game struct {
	status api.GameStatus

	id string

	player1   player
	player2   player
	playerWon player
	players   map[string]mark

	board board
}

type player struct {
	mark mark
	name string
}

type mark = string

func NewGame(playerName string) Game {
	return Game{
		status:    api.GameStatus_WAITING_P2_JOIN,
		id:        strconv.Itoa(time.Now().Nanosecond()),
		player1:   player{mark: "X", name: playerName},
		player2:   player{},
		playerWon: player{},
		players:   map[string]mark{playerName: "X"},
		board:     blankBoard(),
	}
}
