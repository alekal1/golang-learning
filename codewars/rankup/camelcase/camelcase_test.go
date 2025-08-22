package camelcase

import "testing"

func TestCamelCase(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{
			"test case", "TestCase",
		},
		{
			"camel case method", "CamelCaseMethod",
		},
		{
			"say hello ", "SayHello",
		},
		{
			" camel case word", "CamelCaseWord",
		},
		{
			"", "",
		},
	}

	for _, tt := range tests {
		got := CamelCase(tt.input)
		if tt.want != got {
			t.Errorf("'%s' expected, got '%s'", tt.want, got)
		}
	}
}
