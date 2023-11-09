package server

import (
	"context"
	"fmt"
	"log"

	"tictactoe/api"
	"tictactoe/server/game"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GameRepo interface {
	Add(game.Game) error
	GetAll() ([]game.Game, error)
	FindByPlayerName(game.Name) (game.Game, error)
	UpdateByID(game.ID, game.Game) error
	DeleteByID(game.ID) error
}

type service struct {
	gamerepo GameRepo

	api.UnimplementedGameServer
}

func NewService(r GameRepo) *service {
	return &service{gamerepo: r}
}

func (s *service) GetGame(ctx context.Context, req *api.GameRequest) (*api.GameResponse, error) {
	log.Printf("server: get game: %v", req)
	if games, err := s.gamerepo.GetAll(); err == nil {
		fmt.Printf("games: %+v\n", games)
	}
	game, _ := s.gamerepo.FindByPlayerName(game.Name(req.PlayerName))
	// if err != nil {
	// 	return &api.GameResponse{}, errors.Wrap(err, "get game")
	// }
	if game.ID == "" {
		return &api.GameResponse{Status: api.GameStatus_NOT_STARTED}, nil
	}
	fmt.Printf("players: %+v\n", game.Players)
	return makeGameResp(game), nil
}

func (s *service) StartGame(ctx context.Context, req *api.GameRequest) (*api.EmptyResponse, error) {
	log.Printf("server: start game: %v", req)
	if allGames, err := s.gamerepo.GetAll(); err == nil {
		fmt.Printf("games: %+v\n", allGames)
	}

	gam, _ := s.gamerepo.FindByPlayerName(game.Name(req.PlayerName))

	// TODO:
	// if err != nil {
	// 	return &api.EmptyResponse{}, errors.Wrap(err, "start game")
	// }

	// First, try to end the game
	if gam.ID != "" {
		if !gam.IsEnded() {
			return &api.EmptyResponse{}, nil
		}
		s.gamerepo.DeleteByID(gam.ID)

		newGame := game.MakeGame(game.Name(req.PlayerName))
		s.gamerepo.Add(newGame)
		return &api.EmptyResponse{}, nil
	}

	// Find a game without second player
	gm, _ := s.gamerepo.FindByPlayerName("")

	// TODO:
	// if err != nil {
	// 	return &api.EmptyResponse{}, nil
	// }

	// Otherwise, add 2nd player
	if gm.ID != "" && gm.Player2.Name == "" {
		g := gm.SetStatus(api.GameStatus_WAITING_P1_TO_TURN)
		g.Player2 = game.MakePlayer2(req)
		g.Players[game.Name(req.PlayerName)] = "O"

		s.gamerepo.UpdateByID(g.ID, g)
		return &api.EmptyResponse{}, nil
	}
	newGame := game.MakeGame(game.Name(req.PlayerName))
	s.gamerepo.Add(newGame)
	return &api.EmptyResponse{}, nil
}

// TODO: split, too long
func (s *service) Turn(ctx context.Context, req *api.TurnRequest) (*api.EmptyResponse, error) {
	log.Printf("server: start game: %v", req)
	if allGames, err := s.gamerepo.GetAll(); err == nil {
		fmt.Printf("games: %+v\n", allGames)
	}

	gam, err := s.gamerepo.FindByPlayerName(game.Name(req.PlayerName))
	if err != nil {
		return nil, err
	}
	if gam.ID == "" {
		return nil, status.Error(codes.NotFound, "Player has no game")
	}

	turn := game.Key(req.Turn)
	if !turn.IsKey() {
		return &api.EmptyResponse{}, nil
	}
	cel := turn.ToCell()
	if isFilled, err := gam.Board.IsFilled(cel); err != nil {
		if err != nil {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
		if isFilled {
			return &api.EmptyResponse{}, nil
		}
	}

	// TODO: may not be found
	mark := gam.Players[game.Name(req.PlayerName)]
	b, err := game.SetCell(gam.Board, cel, mark)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	gam.Board = b

	s.gamerepo.UpdateByID(gam.ID, gam)

	// Ending the game
	if gam.Board.IsWinner(mark) {
		g := gam.SetStatus(api.GameStatus_WON)
		g.PlayerWon = game.Player{ // todord
			Mark: mark,
			Name: game.Name(req.PlayerName),
		}

		s.gamerepo.UpdateByID(g.ID, g)
		return &api.EmptyResponse{}, nil
	}
	if !gam.Board.HasEmpty() {
		g := gam.SetStatus(api.GameStatus_DRAW)

		s.gamerepo.UpdateByID(g.ID, g)
		return &api.EmptyResponse{}, nil
	}

	// Waiting for turns
	if gam.Player1.Name == game.Name(req.PlayerName) {
		g := gam.SetStatus(api.GameStatus_WAITING_P2_TO_TURN)

		s.gamerepo.UpdateByID(g.ID, g)
		return &api.EmptyResponse{}, nil
	}
	if gam.Player2.Name == game.Name(req.PlayerName) {
		g := gam.SetStatus(api.GameStatus_WAITING_P1_TO_TURN)

		s.gamerepo.UpdateByID(g.ID, g)
		return &api.EmptyResponse{}, nil
	}
	return &api.EmptyResponse{}, nil
}
