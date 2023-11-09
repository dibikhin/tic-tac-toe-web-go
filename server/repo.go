package server

import (
	"tictactoe/api"

	"tictactoe/server/game"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type gameRepo struct {
	gamesDB []game.Game
}

func MakeGameRepo(games ...game.Game) *gameRepo {
	if len(games) == 0 {
		return &gameRepo{
			gamesDB: nil,
		}
	}
	return &gameRepo{
		gamesDB: games,
	}
}

func (r *gameRepo) Add(g game.Game) error {
	r.gamesDB = append(r.gamesDB, g)
	return nil
}

func (r *gameRepo) GetAll() ([]game.Game, error) {
	return r.gamesDB, nil
}

func (r *gameRepo) FindByPlayerName(name game.Name) (game.Game, error) {
	for _, g := range r.gamesDB {
		if (g.Player1.Name == name || g.Player2.Name == name) &&
			g.Status != api.GameStatus_DELETED {
			return g, nil
		}
	}
	return game.Game{}, status.Error(codes.NotFound, "game not found")
}

func (r *gameRepo) UpdateByID(id game.ID, diff game.Game) error {
	for i := range r.gamesDB {
		g := r.gamesDB[i]
		if g.ID == id &&
			g.Status != api.GameStatus_DELETED {
			r.gamesDB[i] = diff

			break
		}
	}
	return nil
}

func (r *gameRepo) DeleteByID(id game.ID) error {
	for i := range r.gamesDB {
		g := r.gamesDB[i]
		if g.ID == id {
			g.Status = api.GameStatus_DELETED
			r.gamesDB[i] = g

			break
		}
	}
	return nil
}
