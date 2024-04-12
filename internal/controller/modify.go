package controller

import (
	"encoding/json"
	"gosha/internal/storage"
	"log"
)

func modify(id string, name string) {
	store, err := storage.GetStore()
	if err != nil {
		log.Fatalf(err.Error())
	}

	tasks, err := storage.Content(store)
	if err != nil {
		log.Fatalf(err.Error())
	}

	var foundTaskIdx = -1
	for idx, task := range tasks {
		if task.Id == id {
			foundTaskIdx = idx
		}
	}

	if foundTaskIdx == -1 {
		log.Fatalf("No such task with id: %s", id)
	}

	// Modify current task
	tasks[foundTaskIdx].Name = name

	// Clear all data from storage
	store.Truncate(0)

	content, err := json.Marshal(tasks)

	if err != nil {
		panic(err)
	}

	// Write new content to the file
	if _, err := store.Write(content); err != nil {
		panic(err)
	}
}
