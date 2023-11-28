package server

import (
	"context"
	"fmt"
	"log"

	"tictactoe/api"
	"tictactoe/server/domain"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GameRepo interface {
	Add(domain.Game) error
	GetAll() ([]domain.Game, error)
	FindByPlayerName(domain.PlayerName) (domain.Game, error)
	UpdateByID(domain.ID, domain.Game) error
	DeleteByID(domain.ID) error
}

type service struct {
	games GameRepo

	api.UnimplementedGameServer
}

func NewGameService(r GameRepo) *service {
	return &service{games: r}
}

func (s *service) GetGame(ctx context.Context, req *api.GameRequest) (*api.GameResponse, error) {
	log.Printf("server: get game: %v", req)
	if games, err := s.games.GetAll(); err == nil {
		fmt.Printf("games: %+v\n", games)
	}
	game, _ := s.games.FindByPlayerName(domain.PlayerName(req.PlayerName))

	if game.ID == "" {
		return &api.GameResponse{Status: api.GameStatus_NOT_STARTED}, nil
	}
	fmt.Printf("players: %+v\n", game.Players)
	return makeGameResp(game), nil
}

func (s *service) StartGame(ctx context.Context, req *api.GameRequest) (*api.EmptyResponse, error) {
	log.Printf("server: start game: %v", req)

	if allGames, err := s.games.GetAll(); err == nil {
		fmt.Printf("games: %+v\n", allGames)
	}

	game, _ := s.games.FindByPlayerName(domain.PlayerName(req.PlayerName))

	// First, try to end the game
	if game.ID != "" {
		if !game.IsEnded() {
			return &api.EmptyResponse{}, nil
		}
		if err := s.games.DeleteByID(game.ID); err != nil {
			return nil, err
		}

		newGame := domain.MakeGame(domain.PlayerName(req.PlayerName))
		if err := s.games.Add(newGame); err != nil {
			return nil, err
		}
		return &api.EmptyResponse{}, nil
	}

	// Find a game without 2nd player
	gm, _ := s.games.FindByPlayerName("") // Due to a game w/ empty Player's name

	// Otherwise, add 2nd player
	if gm.ID != "" && gm.Player2.Name == "" {
		g := gm.WithStatus(api.GameStatus_WAITING_P1_TO_TURN)
		g = g.WithPlayer2(domain.NewPlayer2(req))

		g.Players[domain.PlayerName(req.PlayerName)] = domain.O

		if err := s.games.UpdateByID(g.ID, g); err != nil {
			return nil, err
		}
		return &api.EmptyResponse{}, nil
	}
	newGame := domain.MakeGame(domain.PlayerName(req.PlayerName))
	if err := s.games.Add(newGame); err != nil {
		return nil, err
	}
	return &api.EmptyResponse{}, nil
}

// TODO: split, too long
func (s *service) Turn(ctx context.Context, req *api.TurnRequest) (*api.EmptyResponse, error) {
	log.Printf("server: start game: %v", req)
	if allGames, err := s.games.GetAll(); err == nil {
		fmt.Printf("games: %+v\n", allGames)
	}

	game, err := s.games.FindByPlayerName(domain.PlayerName(req.PlayerName))
	if err != nil {
		return nil, err
	}
	if game.ID == "" {
		return nil, status.Error(codes.NotFound, "Player has no game")
	}

	turn := domain.Key(req.Turn)
	if !turn.IsKey() {
		return &api.EmptyResponse{}, nil
	}
	cel := turn.ToCell()
	if isFilled, err := game.Board.IsFilled(cel); err != nil {
		if err != nil || isFilled {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
	}

	// TODO: may not be found
	mark := game.Players[domain.PlayerName(req.PlayerName)]
	b, err := game.Board.WithCell(cel, mark)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	game = game.WithBoard(b)
	if err := s.games.UpdateByID(game.ID, game); err != nil {
		return nil, err
	}

	// Ending the game
	if game.Board.IsWinner(mark) {
		g := game.WithStatus(api.GameStatus_WON)

		p := domain.Player{
			Mark: mark,
			Name: domain.PlayerName(req.PlayerName),
		}
		g = g.WithPlayerWon(p)
		if err := s.games.UpdateByID(g.ID, g); err != nil {
			return nil, err
		}
		return &api.EmptyResponse{}, nil
	}
	if !game.Board.HasEmpty() {
		g := game.WithStatus(api.GameStatus_DRAW)

		if err := s.games.UpdateByID(g.ID, g); err != nil {
			return nil, err
		}
		return &api.EmptyResponse{}, nil
	}

	// Waiting for turns
	if game.Player1.Name == domain.PlayerName(req.PlayerName) {
		g := game.WithStatus(api.GameStatus_WAITING_P2_TO_TURN)

		if err := s.games.UpdateByID(g.ID, g); err != nil {
			return nil, err
		}
		return &api.EmptyResponse{}, nil
	}
	if game.Player2.Name == domain.PlayerName(req.PlayerName) {
		g := game.WithStatus(api.GameStatus_WAITING_P1_TO_TURN)

		if err := s.games.UpdateByID(g.ID, g); err != nil {
			return nil, err
		}
		return &api.EmptyResponse{}, nil
	}
	return &api.EmptyResponse{}, nil
}
