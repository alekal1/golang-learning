package main

import (
	"testing"
)

func TestReversedString(t *testing.T) {
	original := "world"
	expected := "dlrow"

	result := ReverseString(original)

	if expected != result {
		t.Errorf("'%s' expected, got '%s'", result, expected)
	}
}
