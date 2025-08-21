package hero

import "testing"

func HeroSurvive(t *testing.T) {
	tests := []struct {
		bullets int
		dragons int
		want    bool
	}{
		{10, 5, true},
		{7, 4, false},
		{4, 5, false},
		{100, 40, true},
		{1500, 751, false},
		{0, 1, false},
	}

	for _, tt := range tests {
		got := Hero(tt.bullets, tt.dragons)
		if got != tt.want {
			t.Errorf("'%v' expected, got '%v'", tt.want, got)
		}
	}
}
