package controller

import (
	"encoding/json"
	"fmt"
	"gosha/internal/storage"
	"log"
	"slices"
)

func remove(id string) {
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
		panic(fmt.Sprintf("No such task with id: %s", id))
	}

	tasks = slices.Concat(tasks[:foundTaskIdx], tasks[foundTaskIdx+1:])

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
