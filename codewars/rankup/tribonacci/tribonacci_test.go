package tribonacci

import "testing"

func TestTribonnaci(t *testing.T) {
	tests := []struct {
		signature [3]float64
		num       int
		want      []float64
	}{
		{
			[3]float64{1, 1, 1}, 10, []float64{1, 1, 1, 3, 5, 9, 17, 31, 57, 105},
		},
		{
			[3]float64{0, 0, 1}, 10, []float64{0, 0, 1, 1, 2, 4, 7, 13, 24, 44},
		},
		{
			[3]float64{0, 0, 0}, 10, []float64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
		{
			[3]float64{0.5, 0.5, 0.5}, 30, []float64{0.5, 0.5, 0.5, 1.5, 2.5, 4.5, 8.5, 15.5, 28.5, 52.5, 96.5, 177.5, 326.5, 600.5, 1104.5, 2031.5, 3736.5, 6872.5, 12640.5, 23249.5, 42762.5, 78652.5, 144664.5, 266079.5, 489396.5, 900140.5, 1655616.5, 3045153.5, 5600910.5, 10301680.5},
		},
	}

	for _, tt := range tests {
		got := Tribonacci(tt.signature, tt.num)
		if !arraysEquals(tt.want, got) {
			t.Errorf("'%v' expected, got '%v'", tt.want, got)
		}
	}
}

func arraysEquals(array1, array2 []float64) bool {
	if len(array1) != len(array2) {
		return false
	}

	for i, v := range array1 {
		if v != array2[i] {
			return false
		}
	}
	return true
}
