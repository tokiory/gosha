package controller

import (
	"github.com/rodaine/table"
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

	t := table.New("ID", "Done?", "Name")
	for _, task := range tasks {
		t.AddRow(task.Id, getDoneText(task.Done), task.Name)
	}
	t.Print()
}
