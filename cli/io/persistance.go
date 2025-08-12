package io

import (
	"aleksale/tasks/json"
	"log"
	"os"
)

const (
	filePath = "tasks.json"
)

func WriteFile[T any](tasks []T) {
	content := json.CreateJsonList(tasks)
	if err := os.WriteFile(filePath, content, 0644); err != nil {
		log.Fatal(err)
	}
}

func ReadFile[T any]() []T {
	dat, err := os.ReadFile(filePath)
	if err != nil {
		if createJsonFile() != nil {
			log.Fatal(err)
		}
		return ReadFile[T]()
	}

	return json.ReadJsonList[T](dat)
}

func createJsonFile() error {
	v, err := os.Create(filePath)
	v.Write([]byte("[]"))

	if err != nil {
		log.Fatal(err)
		return err
	}
	v.Close()
	return err
}
