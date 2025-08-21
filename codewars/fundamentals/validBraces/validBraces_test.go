package validBraces

import "testing"

func TestValidBraces(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{
			"(){}[]", true,
		},
		{
			"([{}])", true,
		},
		{
			"({", false,
		},
		{
			"[(])", false,
		},
		{
			"[({)](]", false,
		},
	}

	for _, tt := range tests {
		got := ValidBraces(tt.input)
		if tt.want != got {
			t.Errorf("'%v' expected, got '%v'", tt.want, got)
		}
	}
}
