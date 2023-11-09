package game

import (
	"testing"
	"tictactoe/api"
)

func TestGame_isEnded(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		game Game
		want bool
	}{
		{
			name: "not ended yet",
			game: Game{
				Status: api.GameStatus_WAITING_P2_JOIN,
			},
			want: false,
		},
		{
			name: "unknown status",
			game: Game{
				Status: 123456789,
			},
			want: false,
		},
		{
			name: "someone won",
			game: Game{
				Status: api.GameStatus_WON,
			},
			want: true,
		},
		{
			name: "draw",
			game: Game{
				Status: api.GameStatus_DRAW,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := tt.game.IsEnded(); got != tt.want {
				t.Errorf("Game.IsEnded() = %v, want %v", got, tt.want)
			}
		})
	}
}
