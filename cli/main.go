package main

import (
	"aleksale/tasks/task"
	"fmt"
	"log"
	"os"
	"strconv"
)

const (
	add    = "add"
	delete = "delete"
	list   = "list"
	done   = "done"
)

func event(arguments []string) {
	action := arguments[0]

	switch action {
	case add:
		task.AddTask(task.Task{
			Text: arguments[1],
		})
	case delete:
		v, _ := strconv.Atoi(arguments[1])
		if ok, err := task.DeleteTask(v); err != nil && !ok {
			log.Fatal(err)
		}
	case list:
		fmt.Println(task.ListAll())
	case done:
		v, _ := strconv.Atoi(arguments[1])
		if ok, err := task.Done(v); err != nil && !ok {
			log.Fatal(err)
		}
	default:
		fmt.Println("Uknown commands - " + arguments[0])
	}
}

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println("Usage: go run main.go [add|delete|done|list] [arguments]")
		return
	}

	event(args)
}
