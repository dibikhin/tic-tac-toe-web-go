package gameserver

import (
	"reflect"
	"testing"
)

// Do test because client shows this to player
func TestBoard_String(t *testing.T) {
	tests := []struct {
		name string
		b    Board
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.b.String(); got != tt.want {
				t.Errorf("Board.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBoard_hasEmpty(t *testing.T) {
	tests := []struct {
		name string
		b    Board
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.b.hasEmpty(); got != tt.want {
				t.Errorf("Board.hasEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBoard_isWinner(t *testing.T) {
	type args struct {
		m Mark
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
			if got := tt.b.isWinner(tt.args.m); got != tt.want {
				t.Errorf("Board.isWinner() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_setCell(t *testing.T) {
	type args struct {
		b Board
		c Cell
		m Mark
	}
	tests := []struct {
		name    string
		args    args
		want    Board
		wantErr bool
	}{
		{
			name: "in range",
			args: args{
				b: blankBoard(),
				c: Cell{
					row: 1,
					col: 2,
				},
				m: "X",
			},
			want: Board{
				{__, __, __},
				{__, __, "X"},
				{__, __, __},
			},
			wantErr: false,
		},
		{
			name: "out of range",
			args: args{
				b: blankBoard(),
				c: Cell{
					row: 123,
					col: 321,
				},
				m: "X",
			},
			want: Board{
				{__, __, __},
				{__, __, __},
				{__, __, __},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := setCell(tt.args.b, tt.args.c, tt.args.m)
			if (err != nil) != tt.wantErr {
				t.Errorf("setCell() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("setCell() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBoard_isFilled(t *testing.T) {
	type args struct {
		c Cell
	}
	tests := []struct {
		name    string
		b       Board
		args    args
		want    bool
		wantErr bool
	}{
		{
			name: "empty board",
			b:    blankBoard(),
			args: args{
				c: Cell{
					row: 2, col: 0,
				},
			},
			want:    false,
			wantErr: false,
		},
		{
			name: "is filled",
			b: Board{
				{__, __, __},
				{__, __, __},
				{"X", __, __},
			},
			args: args{
				c: Cell{
					row: 2, col: 0,
				},
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "out of range",
			b:    blankBoard(),
			args: args{
				c: Cell{
					row: 234, col: 543,
				},
			},
			want:    false,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.b.isFilled(tt.args.c)
			if (err != nil) != tt.wantErr {
				t.Errorf("Board.isFilled() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Board.isFilled() = %v, want %v", got, tt.want)
			}
		})
	}
}
