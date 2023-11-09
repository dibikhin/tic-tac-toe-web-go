package game

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
	Players   map[Name]Mark

	Board Board
}

type ID string

type Player struct {
	Mark Mark
	Name Name
}

type Mark string
type Name string

func MakePlayer2(req *api.GameRequest) Player {
	return Player{
		Mark: "O",
		Name: Name(req.PlayerName),
	}
}

func MakeGame(name Name) Game {
	return Game{
		Status:    api.GameStatus_WAITING_P2_JOIN,
		ID:        genID(),
		Player1:   Player{Mark: "X", Name: name}, // Should have at least first player
		Player2:   Player{},
		PlayerWon: Player{},
		Players:   map[Name]Mark{name: "X"},
		Board:     blankBoard(),
	}
}

func genID() ID {
	return ID(strconv.Itoa(time.Now().Nanosecond()))
}

func (g Game) IsEnded() bool {
	return g.Status == api.GameStatus_DRAW ||
		g.Status == api.GameStatus_WON
}

func (g Game) SetStatus(status api.GameStatus) Game {
	g.Status = status
	return g
}
