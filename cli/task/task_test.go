package task

import (
	"os"
	"testing"
)

func deleteJsonFile() {
	os.Remove("tasks.json")
}

func TestAddAndListTask(t *testing.T) {
	const taskText = "Test 1"

	task := Task{
		Text: taskText,
	}

	AddTask(task)
	tasks := ListAll()

	if len(tasks) != 1 {
		t.Fatalf("Expected 1 task, got %d", len(tasks))
	}

	if tasks[0].Text != taskText {
		t.Errorf("Expected task text to be '%s', got '%s'", taskText, tasks[0].Text)
	}

	deleteJsonFile()
}

func TestDeleteTask(t *testing.T) {
	AddTask(Task{Text: "Test file 1"})
	ok, err := DeleteTask(0)

	if !ok || err != nil {
		t.Fatalf("Expected task to be deleted, got error %v", err)
	}

	tasks := ListAll()
	if len(tasks) != 0 {
		t.Errorf("Expected task list to be empty after deletion")
	}

	deleteJsonFile()

}

func TestDeleteTask_NotFound(t *testing.T) {
	ok, err := DeleteTask(1)
	if ok || err == nil {
		t.Errorf("Expected error when deleting non-existent task")
	}

	deleteJsonFile()
}

func TestDone(t *testing.T) {
	AddTask(Task{
		Id:   0,
		Text: "Test task",
	})

	ok, err := Done(0)

	if !ok || err != nil {
		t.Fatalf("Expected task to be completed, got error %v", err)
	}

	tasks := ListAll()
	if !tasks[0].Complete {
		t.Errorf("Expected task to be complete")
	}

	deleteJsonFile()
}

func TestDone_NotFound(t *testing.T) {
	ok, err := Done(1)
	if ok || err == nil {
		t.Errorf("Expected error when completing non-existent task")
	}

	deleteJsonFile()
}
