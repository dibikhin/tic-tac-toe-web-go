package server

import (
	"context"
	"fmt"
	"log"
	"tictactoe/pkg/api"
)

type GameRepo interface {
	Add(Game)
	GetAll() []Game
	FindByPlayerName(string) Game
	UpdateById(string, Game)
	DeleteById(string)
}

type gameService struct {
	gameRepo GameRepo

	api.UnimplementedGameServer
}

func NewGameService(gs GameRepo) *gameService {
	return &gameService{gameRepo: gs}
}

func (s *gameService) GetGame(ctx context.Context, sr *api.GameRequest) (*api.GameResponse, error) {
	log.Printf("srv: get game %v", sr)
	fmt.Printf("games: %+v\n", s.gameRepo.GetAll())

	g := s.gameRepo.FindByPlayerName(sr.PlayerName)
	if g.id == "" {
		return &api.GameResponse{Status: api.GameStatus_NOT_STARTED}, nil
	}
	fmt.Printf("players: %+v\n", g.players)
	return newGameResp(g), nil
}

func newGameResp(g Game) *api.GameResponse {
	return &api.GameResponse{
		Status:    api.GameStatus(g.status),
		Player1:   &api.Player{Mark: g.player1.mark, Name: g.player1.name},
		Player2:   &api.Player{Mark: g.player2.mark, Name: g.player2.name},
		PlayerWon: &api.Player{Mark: g.playerWon.mark, Name: g.playerWon.name},
		Board:     g.board.String(),
	}
}

func (s *gameService) StartGame(ctx context.Context, sr *api.GameRequest) (*api.EmptyResponse, error) {
	log.Printf("srv: start game %v", sr)
	fmt.Printf("games: %+v\n", s.gameRepo.GetAll())

	g := s.gameRepo.FindByPlayerName(sr.PlayerName)
	if g.id != "" {
		if g.status != api.GameStatus_WON && g.status != api.GameStatus_DRAW {
			return &api.EmptyResponse{}, nil
		}
		s.gameRepo.DeleteById(g.id)

		newGame := NewGame(sr.PlayerName)
		s.gameRepo.Add(newGame)
		return &api.EmptyResponse{}, nil
	}
	gg := s.gameRepo.FindByPlayerName("")
	if gg.id != "" && gg.player2.name == "" {
		gg.status = api.GameStatus_WAITING_P1_TO_TURN
		gg.player2 = player{mark: "O", name: sr.PlayerName}
		gg.players[sr.PlayerName] = "O"
		s.gameRepo.UpdateById(gg.id, gg)
		return &api.EmptyResponse{}, nil
	}
	newGame := NewGame(sr.PlayerName)
	s.gameRepo.Add(newGame)
	return &api.EmptyResponse{}, nil
}

func (s *gameService) Turn(ctx context.Context, tr *api.TurnRequest) (*api.EmptyResponse, error) {
	g := s.gameRepo.FindByPlayerName(tr.PlayerName)
	if g.id == "" {
		return nil, fmt.Errorf("Player has no game")
	}
	turn := key(tr.Turn)
	if !turn.isKey() {
		return &api.EmptyResponse{}, nil
	}
	cel := turn.toCell()
	if g.board.isFilled(cel) {
		return &api.EmptyResponse{}, nil
	}
	mark := g.players[tr.PlayerName]
	g.board = setCell(g.board, cel, mark)
	s.gameRepo.UpdateById(g.id, g)
	if g.board.isWinner(mark) {
		g.status = api.GameStatus_WON
		g.playerWon = player{mark, tr.PlayerName}
		s.gameRepo.UpdateById(g.id, g)
		return &api.EmptyResponse{}, nil
	}
	if !g.board.hasEmpty() {
		g.status = api.GameStatus_DRAW
		s.gameRepo.UpdateById(g.id, g)
		return &api.EmptyResponse{}, nil
	}
	if g.player1.name == tr.PlayerName {
		g.status = api.GameStatus_WAITING_P2_TO_TURN
		s.gameRepo.UpdateById(g.id, g)
		return &api.EmptyResponse{}, nil
	}
	if g.player2.name == tr.PlayerName {
		g.status = api.GameStatus_WAITING_P1_TO_TURN
		s.gameRepo.UpdateById(g.id, g)
		return &api.EmptyResponse{}, nil
	}
	return &api.EmptyResponse{}, nil
}
