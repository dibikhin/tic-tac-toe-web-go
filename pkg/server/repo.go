package server

import "tictactoe/pkg/api"

type gameRepo struct {
	games []Game
}

func NewGameRepo() *gameRepo {
	return &gameRepo{
		games: []Game{},
	}
}

func (r *gameRepo) Add(g Game) {
	r.games = append(r.games, g)
}

func (r *gameRepo) GetAll() []Game {
	return r.games
}

func (r *gameRepo) FindByPlayerName(n string) Game {
	for _, g := range r.games {
		if (g.player1.name == n || g.player2.name == n) &&
			g.status != api.GameStatus_DELETED {
			return g
		}
	}
	return Game{}
}

func (r *gameRepo) UpdateById(id string, diff Game) {
	for i := range r.games {
		if r.games[i].id == id &&
			r.games[i].status != api.GameStatus_DELETED {
			r.games[i] = diff
		}
	}
}

func (r *gameRepo) DeleteById(id string) {
	for i := range r.games {
		g := &r.games[i]
		if g.id == id {
			g.status = api.GameStatus_DELETED
		}
	}
}
