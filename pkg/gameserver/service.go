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
	Add(Game) error
	GetAll() ([]Game, error)
	FindByPlayerName(Name) (Game, error)
	UpdateByID(ID, Game) error
	DeleteByID(ID) error
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
	game, _ := s.gamerepo.FindByPlayerName(Name(req.PlayerName))
	// if err != nil {
	// 	return &api.GameResponse{}, errors.Wrap(err, "get game")
	// }
	if game.id == "" {
		return &api.GameResponse{Status: api.GameStatus_NOT_STARTED}, nil
	}
	fmt.Printf("players: %+v\n", game.players)
	return makeGameResp(game), nil
}

func newApiPlayer(p Player) *api.Player {
	return &api.Player{
		Mark: string(p.mark),
		Name: string(p.name),
	}
}

func makeGameResp(g Game) *api.GameResponse {
	return &api.GameResponse{
		Status:    api.GameStatus(g.status),
		Player1:   newApiPlayer(g.player1),
		Player2:   newApiPlayer(g.player2),
		PlayerWon: newApiPlayer(g.playerWon),
		Board:     g.board.String(),
	}
}

func (s *service) StartGame(ctx context.Context, req *api.GameRequest) (*api.EmptyResponse, error) {
	log.Printf("server: start game: %v", req)
	if allGames, err := s.gamerepo.GetAll(); err == nil {
		fmt.Printf("games: %+v\n", allGames)
	}

	game, _ := s.gamerepo.FindByPlayerName(Name(req.PlayerName))

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

		newGame := MakeGame(Name(req.PlayerName))
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
		gg.player2 = Player{mark: "O", name: Name(req.PlayerName)}
		gg.players[Name(req.PlayerName)] = "O"

		s.gamerepo.UpdateByID(gg.id, gg)
		return &api.EmptyResponse{}, nil
	}
	newGame := MakeGame(Name(req.PlayerName))
	s.gamerepo.Add(newGame)
	return &api.EmptyResponse{}, nil
}

func (s *service) Turn(ctx context.Context, req *api.TurnRequest) (*api.EmptyResponse, error) {
	log.Printf("server: start game: %v", req)
	if allGames, err := s.gamerepo.GetAll(); err == nil {
		fmt.Printf("games: %+v\n", allGames)
	}
	game, err := s.gamerepo.FindByPlayerName(Name(req.PlayerName))
	if err != nil {
		return nil, err
	}
	if game.id == "" {
		return nil, status.Error(codes.NotFound, "Player has no game")
	}
	turn := Key(req.Turn)
	if !turn.isKey() {
		return &api.EmptyResponse{}, nil
	}
	cel := turn.toCell()
	if game.board.isFilled(cel) {
		return &api.EmptyResponse{}, nil
	}
	mark := game.players[Name(req.PlayerName)]
	game.board = setCell(game.board, cel, mark)

	s.gamerepo.UpdateByID(game.id, game)

	// Ending the game
	if game.board.isWinner(mark) {
		game.status = api.GameStatus_WON
		game.playerWon = Player{mark, Name(req.PlayerName)}

		s.gamerepo.UpdateByID(game.id, game)
		return &api.EmptyResponse{}, nil
	}
	if !game.board.hasEmpty() {
		game.status = api.GameStatus_DRAW

		s.gamerepo.UpdateByID(game.id, game)
		return &api.EmptyResponse{}, nil
	}

	// Waiting for turns
	if game.player1.name == Name(req.PlayerName) {
		game.status = api.GameStatus_WAITING_P2_TO_TURN

		s.gamerepo.UpdateByID(game.id, game)
		return &api.EmptyResponse{}, nil
	}
	if game.player2.name == Name(req.PlayerName) {
		game.status = api.GameStatus_WAITING_P1_TO_TURN

		s.gamerepo.UpdateByID(game.id, game)
		return &api.EmptyResponse{}, nil
	}
	return &api.EmptyResponse{}, nil
}
