package loveflowers

import "testing"

func TestLoveFlowers(t *testing.T) {
	tests := []struct {
		flower1 int
		flower2 int
		want    bool
	}{
		{1, 4, true},
		{2, 2, false},
		{0, 1, true},
		{100, 101, true},
	}

	for _, tt := range tests {
		got := LoveFunc(tt.flower1, tt.flower2)
		if got != tt.want {
			t.Errorf("'%v' expected, got '%v'", tt.want, got)
		}
	}
}
