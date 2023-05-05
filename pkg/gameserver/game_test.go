package gameserver

import (
	"testing"
	"tictactoe/pkg/api"
)

func TestGame_isEnded(t *testing.T) {
	type fields struct {
		status    api.GameStatus
		id        ID
		player1   Player
		player2   Player
		playerWon Player
		players   map[Name]Mark
		board     Board
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := Game{
				status:    tt.fields.status,
				id:        tt.fields.id,
				player1:   tt.fields.player1,
				player2:   tt.fields.player2,
				playerWon: tt.fields.playerWon,
				players:   tt.fields.players,
				board:     tt.fields.board,
			}
			if got := g.isEnded(); got != tt.want {
				t.Errorf("Game.isEnded() = %v, want %v", got, tt.want)
			}
		})
	}
}
