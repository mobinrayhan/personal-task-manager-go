package task

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"

	"mobin.dev/io"
)

type ID = string

type Task struct {
	ID          ID        `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	DueDate     time.Time `json:"due_data"`
}

const FILE_NAME = "task_list_.json"

func Add(t Task) {
	tasks := List()
	tasksLen := len(tasks)
	tasks[tasksLen+1] = t
	io.WriteToFile(FILE_NAME, tasks)
}

func List() map[int]Task {
	tasks := map[int]Task{}
	err := io.ReadFromFile(FILE_NAME, &tasks)

	if err != nil {
		fmt.Println("Error loading tasks:", err)
	}

	return tasks
}

func Update(id ID, t Task) (ID, error) {
	tasks := List()
	for i, v := range tasks {
		if id == v.ID {
			t.ID = id
			tasks[i] = t
			break
		}
	}
	jsonDate, _ := json.MarshalIndent(tasks, "", " ")
	err := os.WriteFile(FILE_NAME, jsonDate, 0644)

	if err != nil {
		return id, errors.New("failed to write at json file")
	}

	return id, nil
}

func Delete(id ID) (string, error) {
	tasks := List()
	var itemKey int = -1

	for key, v := range tasks {
		fmt.Println(v, key)
		if v.ID == id {
			itemKey = key
			break
		}
	}

	if itemKey == -1 {
		formattedMsg := fmt.Sprintf("Task item not found ID : %s", id)
		return "", errors.New(formattedMsg)
	}

	delete(tasks, itemKey)
	io.WriteToFile(FILE_NAME, &tasks)

	deleteMessage := fmt.Sprintf("Delete Task Successfully %s:", id)
	return deleteMessage, nil
}
