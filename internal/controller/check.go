package controller

import (
	"encoding/json"
	"fmt"
	"gosha/internal/storage"
)

func check(id string) {
  store, err := storage.GetStore()
  if err != nil {
    panic(err)
  }

  tasks, err := storage.Content(store)
  if err != nil {
    panic(err)
  }


  // Find required task
  var foundTaskIdx = -1
  for idx, task := range tasks {
    if task.Id == id {
      foundTaskIdx = idx
    }
  }

  if foundTaskIdx == -1 {
    panic(fmt.Sprintf("No such task with id: %s", id))
  }

  // Toggle founded task
  tasks[foundTaskIdx].Done = !tasks[foundTaskIdx].Done

  // Truncate old file
  store.Truncate(0)

  content, err := json.Marshal(tasks, )

  if err != nil {
    panic(err)
  }

  // Write new content to the file
  if _, err := store.Write(content); err != nil {
    panic(err)
  }
}
