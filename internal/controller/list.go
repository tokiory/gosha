package controller

import (
	"fmt"
	"gosha/internal/storage"
	"log"
)

func list() {
	store, err := storage.GetStore()
	if err != nil {
		log.Fatalf(err.Error())
	}

	tasks, err := storage.Content(store)
	if err != nil {
		log.Fatalf(err.Error())
	}

	for _, task := range tasks {
		if !task.Done {
			fmt.Println(task.Id, task.Name)
		}
	}
}
