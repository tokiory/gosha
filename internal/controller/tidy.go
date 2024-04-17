package controller

import (
	"bufio"
	"encoding/json"
	"fmt"
	"gosha/internal/storage"
	"gosha/internal/task"
	"os"
	"strings"
)

func tidy() {
	store, err := storage.GetStore()
	if err != nil {
		panic(err)
	}

	tasks, err := storage.Content(store)
	if err != nil {
		panic(err)
	}

	list()

	fmt.Println("Are you sure? (y/n)")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	fmt.Println()
	if formattedInput := strings.Trim(strings.ToLower(input), "\n")[0]; formattedInput != 'y' {
		fmt.Println("It's ok, you can do it later :)")
		os.Exit(0)
	}

	// Filter tasks
	filteredTasks := make([]task.Task, 0, len(tasks))

	for _, task := range tasks {
		if !task.Done {
			filteredTasks = append(filteredTasks, task)
		}
	}

	// Clear all data from storage
	store.Truncate(0)

	// Marshall filtered tasks
	content, err := json.Marshal(filteredTasks)

	if err != nil {
		panic(err)
	}

	// Write new content to the file
	if _, err := store.Write(content); err != nil {
		panic(err)
	}

	fmt.Println("Clear!")
}
