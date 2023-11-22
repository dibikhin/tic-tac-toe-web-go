package server

import (
	"sync"
	"tictactoe/api"

	"tictactoe/server/domain"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type repo struct {
	games []domain.Game
	mu    sync.Mutex
}

func MakeGameRepo(games ...domain.Game) *repo {
	return &repo{
		games: games,
	}
}

func (r *repo) Add(g domain.Game) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.games = append(r.games, g)

	return nil
}

func (r *repo) GetAll() ([]domain.Game, error) {
	return r.games, nil
}

func (r *repo) FindByPlayerName(name domain.PlayerName) (domain.Game, error) {
	for _, g := range r.games {
		if !g.IsDeleted() &&
			(g.Player1.Name == name || g.Player2.Name == name) {
			return g, nil
		}
	}
	return domain.Game{}, status.Error(codes.NotFound, "game not found")
}

func (r *repo) UpdateByID(id domain.ID, update domain.Game) error {
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

func (r *repo) DeleteByID(id domain.ID) error {
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
