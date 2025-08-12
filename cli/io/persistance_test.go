package io

import (
	"os"
	"testing"
)

type TestTask struct {
	Id       int
	Text     string
	Complete bool
}

func TestWriteAndReadFile(t *testing.T) {
	original := []TestTask{
		{Id: 0, Text: "Task 1"},
		{Id: 0, Text: "Task 2"},
	}

	WriteFile(original)

	read := ReadFile[TestTask]()

	if len(read) != len(original) {
		t.Fatalf("Expected %d tasks, got %d", len(original), len(read))
	}

	for i := range read {
		if read[i] != original[i] {
			t.Errorf("Mismatch at index %d: expected %+v, got %+v", i, original[i], read[i])
		}
	}

	os.Remove(filePath)
}
