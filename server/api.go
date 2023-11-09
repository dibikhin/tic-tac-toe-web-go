package server

import (
	"tictactoe/api"
	"tictactoe/server/game"
)

func makeGameResp(g game.Game) *api.GameResponse {
	return &api.GameResponse{
		Status:    api.GameStatus(g.Status),
		Player1:   toAPIPlayer(g.Player1),
		Player2:   toAPIPlayer(g.Player2),
		PlayerWon: toAPIPlayer(g.PlayerWon),
		Board:     g.Board.String(),
	}
}

func toAPIPlayer(p game.Player) *api.Player {
	return &api.Player{
		Mark: string(p.Mark),
		Name: string(p.Name),
	}
}
