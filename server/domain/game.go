package domain

import (
	"strconv"
	"tictactoe/api"
	"time"
)

type Game struct {
	Status api.GameStatus

	ID ID

	Player1   Player
	Player2   Player
	PlayerWon Player
	Players   map[PlayerName]Mark // todo: Players vs P1+P2

	Board Board
}

const X Mark = "X"
const O Mark = "O" // It's a letter, not zero :)

type (
	ID         string
	Mark       string
	PlayerName string

	Player struct {
		Mark Mark
		Name PlayerName
	}
)

func MakePlayer2(req *api.GameRequest) Player {
	return Player{
		Mark: O,
		Name: PlayerName(req.PlayerName),
	}
}

func MakeGame(name PlayerName) Game {
	return Game{
		Status:    api.GameStatus_WAITING_P2_JOIN,
		ID:        genID(),
		Player1:   makePlayer1(name), // Should have at least 1st player
		Player2:   emptyPlayer(),
		PlayerWon: emptyPlayer(),
		Players:   map[PlayerName]Mark{name: X},
		Board:     blankBoard(),
	}
}

func makePlayer1(name PlayerName) Player {
	return Player{Mark: X, Name: name}
}

func emptyPlayer() Player {
	return Player{}
}

func genID() ID {
	return ID(strconv.Itoa(time.Now().Nanosecond()))
}

func (g Game) IsEnded() bool {
	return g.Status == api.GameStatus_DRAW ||
		g.Status == api.GameStatus_WON
}

func (g Game) IsDeleted() bool {
	return g.Status == api.GameStatus_DELETED
}

func (g Game) WithStatus(s api.GameStatus) Game {
	g.Status = s
	return g
}
