package gameserver

import (
	"tictactoe/pkg/api"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type gameRepo struct {
	gamesDB []Game
}

func MakeGameRepo(games ...Game) *gameRepo {
	if len(games) == 0 {
		return &gameRepo{
			gamesDB: nil,
		}
	}
	return &gameRepo{
		gamesDB: games,
	}
}

func (r *gameRepo) Add(g Game) error {
	r.gamesDB = append(r.gamesDB, g)
	return nil
}

func (r *gameRepo) GetAll() ([]Game, error) {
	return r.gamesDB, nil
}

func (r *gameRepo) FindByPlayerName(name Name) (Game, error) {
	for _, g := range r.gamesDB {
		if (g.player1.name == name || g.player2.name == name) &&
			g.status != api.GameStatus_DELETED {
			return g, nil
		}
	}
	return Game{}, status.Error(codes.NotFound, "game not found")
}

func (r *gameRepo) UpdateByID(id ID, diff Game) error {
	for i := range r.gamesDB {
		if r.gamesDB[i].id == id &&
			r.gamesDB[i].status != api.GameStatus_DELETED {
			r.gamesDB[i] = diff

			break
		}
	}
	return nil
}

func (r *gameRepo) DeleteByID(id ID) error {
	for i := range r.gamesDB {
		g := &r.gamesDB[i]
		if g.id == id {
			g.status = api.GameStatus_DELETED
			break
		}
	}
	return nil
}
