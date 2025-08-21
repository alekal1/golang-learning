package smallestInt

import "testing"

func TestSmallestInt(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{
			[]int{34, 15, 88, 2}, 2,
		},
		{
			[]int{34, -345, -1, 100}, -345,
		},
	}

	for _, tt := range tests {
		got := SmallestIntegerFinder(tt.nums)
		if got != tt.want {
			t.Errorf("'%d' expected, got '%d'", tt.want, got)
		}
	}
}
