package controller

import (
	"gosha/internal/storage"
	"gosha/internal/task"
	"encoding/json"
)

func add(name string) {
  store, err := storage.GetStore()
  if err != nil {
    panic(err)
  }

  list, err := storage.Content(store)
  if err != nil {
    panic(err)
  }

  list = append(list, task.New(name))
  content, err := json.Marshal(list)

  if err != nil {
    panic(err)
  }

  // Truncate old file
  store.Truncate(0)

  // Write new content to the file
  if _, err := store.Write(content); err != nil {
    panic(err)
  }
}
