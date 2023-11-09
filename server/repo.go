package server

import (
	"sync"
	"tictactoe/api"

	"tictactoe/server/game"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type gameRepo struct {
	games []game.Game
	mu    sync.Mutex
}

func MakeGameRepo(games ...game.Game) *gameRepo {
	if len(games) == 0 {
		return &gameRepo{
			games: nil,
		}
	}
	return &gameRepo{
		games: games,
	}
}

func (r *gameRepo) Add(g game.Game) error {
	r.mu.Lock()
	r.games = append(r.games, g)
	r.mu.Unlock()

	return nil
}

func (r *gameRepo) GetAll() ([]game.Game, error) {
	return r.games, nil
}

func (r *gameRepo) FindByPlayerName(name game.PlayerName) (game.Game, error) {
	for _, g := range r.games {
		if !g.IsDeleted() &&
			(g.Player1.Name == name || g.Player2.Name == name) {
			return g, nil
		}
	}
	return game.Game{}, status.Error(codes.NotFound, "game not found")
}

func (r *gameRepo) UpdateByID(id game.ID, update game.Game) error {
	for i := range r.games {
		g := r.games[i]
		if g.ID == id && !g.IsDeleted() {
			r.mu.Lock()
			r.games[i] = update
			r.mu.Unlock()

			break
		}
	}
	return nil
}

func (r *gameRepo) DeleteByID(id game.ID) error {
	for i := range r.games {
		gam := r.games[i]
		if gam.ID == id {
			g := gam.WithStatus(api.GameStatus_DELETED)

			r.mu.Lock()
			r.games[i] = g
			r.mu.Unlock()

			break
		}
	}
	return nil
}
