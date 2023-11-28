package domain

import (
	"reflect"
	"testing"
)

// NOTE: Testing the fn because clients show the string to player
func TestBoard_String(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		b    Board
		want string
	}{
		{
			name: "empty board",
			b:    blankBoard(),
			want: "- - -\n- - -\n- - -",
		},
		{
			name: "some cells filled",
			b: Board{
				{"X", "X", "O"},
				{"X", "-", "X"},
				{"O", "O", "-"},
			},
			want: "X X O\nX - X\nO O -",
		},
		{
			name: "full board",
			b: Board{
				{"X", "X", "X"},
				{"X", "X", "X"},
				{"X", "X", "X"},
			},
			want: "X X X\nX X X\nX X X",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := tt.b.String(); got != tt.want {
				t.Errorf("Board.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBoard_hasEmpty(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		b    Board
		want bool
	}{
		{
			name: "all empty",
			b:    blankBoard(),
			want: true,
		},
		{
			name: "some empty",
			b: Board{
				{__, __, "X"},
				{"X", __, __},
				{__, "X", __},
			},
			want: true,
		},
		{
			name: "no empty",
			b: Board{
				{"X", "X", "X"},
				{"X", "X", "X"},
				{"X", "X", "X"},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := tt.b.HasEmpty(); got != tt.want {
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
		{
			name: "empty board",
			b: Board{
				{__, __, __},
				{__, __, __},
				{__, __, __},
			},
			args: args{
				m: "X",
			},
			want: false,
		},
		{
			name: "some marks",
			b: Board{
				{"X", __, __},
				{__, "O", __},
				{__, __, __},
			},
			args: args{
				m: "X",
			},
			want: false,
		},
		{
			name: "X horizontal win",
			b: Board{
				{__, __, __},
				{__, "O", __},
				{"X", "X", "X"},
			},
			args: args{
				m: "X",
			},
			want: true,
		},
		{
			name: "X vertical win",
			b: Board{
				{__, __, "X"},
				{__, "O", "X"},
				{__, __, "X"},
			},
			args: args{
				m: "X",
			},
			want: true,
		},
		{
			name: "X diagonal win",
			b: Board{
				{__, __, "X"},
				{__, "X", "X"},
				{"X", __, __},
			},
			args: args{
				m: "X",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := tt.b.IsWinner(tt.args.m); got != tt.want {
				t.Errorf("Board.isWinner() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_SetCell(t *testing.T) {
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
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := tt.args.b.WithCell(tt.args.c, tt.args.m)
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
	t.Parallel()

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
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := tt.b.IsFilled(tt.args.c)
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
