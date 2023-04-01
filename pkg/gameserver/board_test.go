package gameserver

import "testing"

func Test_board_isFilled(t *testing.T) {
	type args struct {
		c Cell
	}
	tests := []struct {
		name string
		b    Board
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.b.isFilled(tt.args.c); got != tt.want {
				t.Errorf("board.isFilled() = %v, want %v", got, tt.want)
			}
		})
	}
}
