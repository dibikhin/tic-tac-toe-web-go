package gameserver

import (
	"errors"
	"tictactoe/pkg/api"
)

type gameRepo struct {
	gamesDB []game
}

func MakeGameRepo(games ...game) *gameRepo {
	if len(games) == 0 {
		return &gameRepo{
			gamesDB: nil,
		}
	}
	return &gameRepo{
		gamesDB: games,
	}
}

func (r *gameRepo) Add(g game) error {
	r.gamesDB = append(r.gamesDB, g)
	return nil
}

func (r *gameRepo) GetAll() ([]game, error) {
	return r.gamesDB, nil
}

func (r *gameRepo) FindByPlayerName(name name) (game, error) {
	for _, g := range r.gamesDB {
		if (g.player1.name == name || g.player2.name == name) &&
			g.status != api.GameStatus_DELETED {
			return g, nil
		}
	}
	return game{}, errors.New("game not found")
}

func (r *gameRepo) UpdateByID(id string, diff game) error {
	for i := range r.gamesDB {
		if r.gamesDB[i].id == id &&
			r.gamesDB[i].status != api.GameStatus_DELETED {
			r.gamesDB[i] = diff

			break
		}
	}
	return nil
}

func (r *gameRepo) DeleteByID(id string) error {
	for i := range r.gamesDB {
		g := &r.gamesDB[i]
		if g.id == id {
			g.status = api.GameStatus_DELETED
			break
		}
	}
	return nil
}
