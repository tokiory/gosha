package controller

import (
	"github.com/rodaine/table"
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

	t := table.New("ID", "Name")

	for _, task := range tasks {
		if !task.Done {
			t.AddRow(task.Id, task.Name)
		}
	}

	t.Print()
}
