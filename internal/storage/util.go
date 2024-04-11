package storage

import (
	"bytes"
	"encoding/json"
	"gosha/internal/task"
	"os"
)

// Get store file
func GetStore() (*os.File, error) {
  var store *os.File
  var err error

  if isStorageExists, _ := Exist(); isStorageExists {
    store, err = Get()
  } else {
    store, err = Create()
  }

  if err != nil {
    return nil, err
  }

  return store, nil
}

// Get content of the storage
func Content(store *os.File) ([]task.Task, error) {

  // Initialize slice for all tasks
  var list []task.Task

  // Read content of current storage
  buffer := new(bytes.Buffer)
  _, err := buffer.ReadFrom(store)

  if err != nil {
    return list, err
  }

  // Get current json
  json.Unmarshal(buffer.Bytes(), &list)

  return list, nil
}
