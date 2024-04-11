
package controller

import (
	"fmt"
	"gosha/internal/storage"
	"log"
)

func all() {
  store, err := storage.GetStore()
  if err != nil {
    log.Fatalf(err.Error())
  }
  
  tasks, err := storage.Content(store)
  if err != nil {
    log.Fatalf(err.Error())
  }

  for _, task := range tasks {
    fmt.Println(task.Id, getDoneText(task.Done), task.Name)
  }
}