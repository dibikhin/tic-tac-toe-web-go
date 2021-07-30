package internal

import (
	"reflect"
	"testing"
)

func Test_board_isFilled(t *testing.T) {
	tests := []struct {
		name string
		b    Board
		args cell
		want bool
	}{
		{"filled", Board{
			{"X", __, "X"},
			{"O", "X", "O"},
			{"X", __, "O"},
		}, cell{0, 0}, true},
		{"empty", Board{
			{"X", __, "X"},
			{"O", __, "O"},
			{"X", __, "O"},
		}, cell{0, 1}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.b.isFilled(tt.args); got != tt.want {
				t.Errorf("board.isFilled() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_board_hasEmpty(t *testing.T) {
	tests := []struct {
		name string
		b    Board
		want bool
	}{
		{"has empty", Board{
			{"X", __, "X"},
			{"O", __, "O"},
			{"X", __, "O"},
		}, true},
		{"all filled", Board{
			{"X", "O", "X"},
			{"O", "X", "O"},
			{"O", "X", "O"},
		}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.b.hasEmpty(); got != tt.want {
				t.Errorf("board.hasEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_board_isWinner(t *testing.T) {
	tests := []struct {
		name string
		b    Board
		arg  string
		want bool
	}{
		{"first row, X", Board{
			{"X", "X", "X"},
			{"O", __, __},
			{"O", __, __},
		}, "X", true},
		{"last col, O", Board{
			{"X", "X", "O"},
			{__, __, "O"},
			{__, __, "O"},
		}, "O", true},
		{"left diagonal, O", Board{
			{"X", "X", "O"},
			{__, "O", __},
			{"O", __, __},
		}, "O", true},
		{"draw", Board{
			{"X", "O", "O"},
			{"O", "X", "X"},
			{"O", "X", "O"},
		}, "O", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.b.isWinner(tt.arg); got != tt.want {
				t.Errorf("board.isWinner() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_board_setCell(t *testing.T) {
	type args struct {
		c      cell
		player string
	}
	tests := []struct {
		name string
		b    Board
		args args
		want Board
	}{
		{
			"1,1",
			Board{
				{__, __, __},
				{__, __, __},
				{__, __, __},
			},
			args{key("5").toCell(), "X"},
			Board{
				{__, __, __},
				{__, "X", __},
				{__, __, __},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			brd := setCell(tt.b, tt.args.c, tt.args.player)
			if !reflect.DeepEqual(brd, tt.want) {
				t.Errorf("setCell() = %v, want %v", brd, tt.want)
			}
		})
	}
}
