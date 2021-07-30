package internal

import (
	"testing"
)

func Test_arrange(t *testing.T) {
	tests := []struct {
		name  string
		arg   string
		want1 player
		want2 player
	}{
		{"empty", "", player{"O", 1}, player{"X", 2}},
		{"space", "    ", player{"O", 1}, player{"X", 2}},
		{"junk", "^G$FDTeg39dslf*^58)#", player{"O", 1}, player{"X", 2}},
		{"small x", "x", player{"X", 1}, player{"O", 2}},
		{"big X", "X", player{"X", 1}, player{"O", 2}},
		{"small o", "o", player{"O", 1}, player{"X", 2}},
		{"big O", "O", player{"O", 1}, player{"X", 2}},
		{"zero", "0", player{"O", 1}, player{"X", 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got1, got2 := arrangePlayers(tt.arg)
			if got1 != tt.want1 {
				t.Errorf("arrange() got = %v, want %v", got1, tt.want1)
			}
			if got2 != tt.want2 {
				t.Errorf("arrange() got1 = %v, want %v", got2, tt.want2)
			}
		})
	}
}

func Test_isKey(t *testing.T) {
	tests := []struct {
		name string
		arg  key
		want bool
	}{
		{"empty", "", false},
		{"space", "  ", false},
		{"too big", "42", false},
		{"too small", "0", false},
		{"letter", "a", false},
		{"junk", " bsd &% #&#", false},
		{"number 1 to 9 incl.", "5", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.arg.isKey(); got != tt.want {
				t.Errorf("isKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_toCell(t *testing.T) {
	tests := []struct {
		name  string
		arg   key
		want  int
		want1 int
	}{
		{"out of range", "-1", 0, 0},
		{"1", "1", 0, 0},
		{"5", "5", 1, 1},
		{"8", "8", 2, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := tt.arg.toCell()
			if c.row != tt.want {
				t.Errorf("pos() got = %v, want %v", c.row, tt.want)
			}
			if c.col != tt.want1 {
				t.Errorf("pos() got1 = %v, want %v", c.col, tt.want1)
			}
		})
	}
}
