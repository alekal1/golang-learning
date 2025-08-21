package points

import "testing"

func TestPoints(t *testing.T) {
	tests := []struct {
		games []string
		want  int
	}{
		{
			[]string{"1:0", "2:0", "3:0", "4:0", "2:1", "3:1", "4:1", "3:2", "4:2", "4:3"}, 30,
		},
		{
			[]string{"1:1", "2:2", "3:3", "4:4", "2:2", "3:3", "4:4", "3:3", "4:4", "4:4"}, 10,
		},
		{
			[]string{"0:1", "0:2", "0:3", "0:4", "1:2", "1:3", "1:4", "2:3", "2:4", "3:4"}, 0,
		},
		{
			[]string{"1:0", "2:0", "3:0", "4:0", "2:1", "1:3", "1:4", "2:3", "2:4", "3:4"}, 15,
		},
	}

	for _, tt := range tests {
		got := Points(tt.games)
		if got != tt.want {
			t.Errorf("'%d' expected, got '%d'", tt.want, got)
		}
	}
}
