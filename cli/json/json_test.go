package json

import (
	"testing"
)

type TestTask struct {
	Id       int
	Text     string
	Complete bool
}

func TestCreateReadJsonList(t *testing.T) {
	original := []TestTask{
		{Id: 0, Text: "Task1"},
		{Id: 1, Text: "Task2"},
	}

	created := CreateJsonList(original)

	read := ReadJsonList[TestTask](created)

	for i := 0; i < len(read); i++ {
		if original[i] != read[i] {
			t.Errorf("Mismatch at index %d: expected %+v, got %+v", i, original[i], read[i])
		}
	}
}
