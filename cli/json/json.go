package json

import (
	"encoding/json"
	"log"
)

func CreateJsonList[T any](tasks []T) []byte {
	data, err := json.Marshal(tasks)
	if err != nil {
		log.Fatal(err)
	}
	return data
}

func ReadJsonList[T any](content []byte) []T {
	var r []T
	err := json.Unmarshal(content, &r)
	if err != nil {
		log.Fatal(err)
	}
	return r
}
