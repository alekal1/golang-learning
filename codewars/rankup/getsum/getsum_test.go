package getsum

import "testing"

func TestGetSum(t *testing.T) {
	tests := []struct {
		a    int
		b    int
		want int
	}{
		{0, 1, 1},
		{1, 2, 3},
		{5, -1, 14},
		{505, 4, 127759},
		{321, 123, 44178},
		{0, -1, -1},
		{-50, 0, -1275},
		{-1, -5, -15},
		{-5, -5, -5},
		{-505, 4, -127755},
		{-321, 123, -44055},
		{0, 0, 0},
		{-5, -1, -15},
	}

	for _, tt := range tests {
		got := GetSum(tt.a, tt.b)
		if tt.want != got {
			t.Errorf("'%d' got, expected '%d'", got, tt.want)
		}
	}
}
