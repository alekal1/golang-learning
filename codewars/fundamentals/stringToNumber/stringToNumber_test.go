package stringToNumber

import "testing"

func TestNumbers(t *testing.T) {
	numbers := map[string]int{
		"1234": 1234,
		"605":  605,
		"1405": 1405,
		"-7":   -7,
	}

	for k, v := range numbers {
		res := StringToNumber(k)
		if res != v {
			t.Errorf("'%d' expected, got '%d'", v, res)
		}
	}
}
