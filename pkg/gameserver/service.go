package gameserver

import (
	"context"
	"fmt"
	"log"

	"tictactoe/pkg/api"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GameRepo interface {
	Add(game) error
	GetAll() ([]game, error)
	FindByPlayerName(name) (game, error)
	UpdateByID(string, game) error
	DeleteByID(string) error
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
	game, _ := s.gamerepo.FindByPlayerName(req.PlayerName)
	// if err != nil {
	// 	return &api.GameResponse{}, errors.Wrap(err, "get game")
	// }
	if game.id == "" {
		return &api.GameResponse{Status: api.GameStatus_NOT_STARTED}, nil
	}
	fmt.Printf("players: %+v\n", game.players)
	return makeGameResp(game), nil
}

func makeGameResp(g game) *api.GameResponse {
	return &api.GameResponse{
		Status:    api.GameStatus(g.status),
		Player1:   &api.Player{Mark: g.player1.mark, Name: g.player1.name},
		Player2:   &api.Player{Mark: g.player2.mark, Name: g.player2.name},
		PlayerWon: &api.Player{Mark: g.playerWon.mark, Name: g.playerWon.name},
		Board:     g.board.String(),
	}
}

func (s *service) StartGame(ctx context.Context, req *api.GameRequest) (*api.EmptyResponse, error) {
	log.Printf("server: start game: %v", req)
	if allGames, err := s.gamerepo.GetAll(); err == nil {
		fmt.Printf("games: %+v\n", allGames)
	}

	game, _ := s.gamerepo.FindByPlayerName(req.PlayerName)

	// TODO
	// if err != nil {
	// 	return &api.EmptyResponse{}, errors.Wrap(err, "start game")
	// }

	// First, try to end the game
	if game.id != "" {
		if !game.isEnded() {
			return &api.EmptyResponse{}, nil
		}
		s.gamerepo.DeleteByID(game.id)

		newGame := MakeGame(req.PlayerName)
		s.gamerepo.Add(newGame)
		return &api.EmptyResponse{}, nil
	}

	gg, _ := s.gamerepo.FindByPlayerName("")

	// TODO
	// if err != nil {
	// 	return &api.EmptyResponse{}, nil
	// }

	// Otherwise, add 2nd player
	if gg.id != "" && gg.player2.name == "" {
		gg.status = api.GameStatus_WAITING_P1_TO_TURN
		gg.player2 = player{mark: "O", name: req.PlayerName}
		gg.players[req.PlayerName] = "O"

		s.gamerepo.UpdateByID(gg.id, gg)
		return &api.EmptyResponse{}, nil
	}
	newGame := MakeGame(req.PlayerName)
	s.gamerepo.Add(newGame)
	return &api.EmptyResponse{}, nil
}

func (s *service) Turn(ctx context.Context, req *api.TurnRequest) (*api.EmptyResponse, error) {
	log.Printf("server: start game: %v", req)
	if allGames, err := s.gamerepo.GetAll(); err == nil {
		fmt.Printf("games: %+v\n", allGames)
	}
	game, err := s.gamerepo.FindByPlayerName(req.PlayerName)
	if err != nil {
		return nil, err
	}
	if game.id == "" {
		return nil, status.Errorf(codes.FailedPrecondition, "Player has no game")
	}
	turn := key(req.Turn)
	if !turn.isKey() {
		return &api.EmptyResponse{}, nil
	}
	cel := turn.toCell()
	if game.board.isFilled(cel) {
		return &api.EmptyResponse{}, nil
	}
	mark := game.players[req.PlayerName]
	game.board = setCell(game.board, cel, mark)

	s.gamerepo.UpdateByID(game.id, game)

	// Ending the game
	if game.board.isWinner(mark) {
		game.status = api.GameStatus_WON
		game.playerWon = player{mark, req.PlayerName}

		s.gamerepo.UpdateByID(game.id, game)
		return &api.EmptyResponse{}, nil
	}
	if !game.board.hasEmpty() {
		game.status = api.GameStatus_DRAW

		s.gamerepo.UpdateByID(game.id, game)
		return &api.EmptyResponse{}, nil
	}

	// Waiting for turns
	if game.player1.name == req.PlayerName {
		game.status = api.GameStatus_WAITING_P2_TO_TURN

		s.gamerepo.UpdateByID(game.id, game)
		return &api.EmptyResponse{}, nil
	}
	if game.player2.name == req.PlayerName {
		game.status = api.GameStatus_WAITING_P1_TO_TURN

		s.gamerepo.UpdateByID(game.id, game)
		return &api.EmptyResponse{}, nil
	}
	return &api.EmptyResponse{}, nil
}
