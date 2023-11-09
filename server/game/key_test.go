package game

import "testing"

func TestKey_isKey(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		k    Key
		want bool
	}{
		{
			name: "a key",
			k:    Key("5"),
			want: true,
		},
		{
			name: "not a key one",
			k:    Key("0"),
			want: false,
		},
		{
			name: "not a key two",
			k:    Key("11"),
			want: false,
		},
		{
			name: "cannot convert to int",
			k:    Key("asdf"),
			want: false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := tt.k.IsKey(); got != tt.want {
				t.Errorf("Key.IsKey() = %v, want %v", got, tt.want)
			}
		})
	}
}
