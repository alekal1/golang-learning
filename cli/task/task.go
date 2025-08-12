package task

import (
	"aleksale/tasks/io"
	"fmt"
)

type Task struct {
	Id       int    `json:"id"`
	Text     string `json:"text"`
	Complete bool   `json:"complete"`
}

func (t Task) String() string {
	check := " "
	if t.Complete {
		check = "X"
	}
	return fmt.Sprintf("%v. [%v] %v\n", t.Id, check, t.Text)
}

func ListAll() []Task {
	return io.ReadFile[Task]()
}

func AddTask(t Task) {
	taskList := ListAll()

	taskList = append(taskList, t)
	arrangeIds(taskList)

	io.WriteFile(taskList)
}

func DeleteTask(tIndex int) (bool, error) {
	taskList := ListAll()

	for i, v := range taskList {
		if tIndex == v.Id {
			taskList = append(taskList[:i], taskList[i+1:]...)
			arrangeIds(taskList)
			io.WriteFile(taskList)
			return true, nil
		}
	}
	return false, TaskNotPresent{
		msg: fmt.Sprintf("Cannot delete task with id %d", tIndex),
	}
}

func Done(tIndex int) (bool, error) {
	taskList := ListAll()

	for i := range taskList {
		if taskList[i].Id == tIndex {
			taskList[i].Complete = true

			io.WriteFile(taskList)
			return true, nil
		}
	}
	return false, TaskNotPresent{
		msg: fmt.Sprintf("Task with ID %d not found", tIndex),
	}
}

func arrangeIds(taskList []Task) {
	for i := range taskList {
		taskList[i].Id = i
	}
}

type TaskNotPresent struct {
	msg string
}

func (np TaskNotPresent) Error() string {
	return np.msg
}
