package subarraysum

import "testing"

func TestSubArraySum(t *testing.T) {
	tests := []struct {
		numbers []int
		want    int
	}{
		{
			[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, 6,
		},
		{
			[]int{-2, -1, -3, -4, -1, -2, -1, -5, -4}, 0,
		},
	}

	for _, tt := range tests {
		got := MaximumSubarraySum(tt.numbers)
		if tt.want != got {
			t.Errorf("'%d' expected, got '%d'", tt.want, got)
		}
	}
}
