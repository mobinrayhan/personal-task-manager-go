package task

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"
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

// var taskList := map[]

const FILE_NAME = "task_list_.json"

func Add(t Task) {
	tasks := List()

	tasksLen := len(tasks)

	if len(tasks) > 0 {
		tasks[tasksLen+1] = t
	} else {
		tasks[tasksLen+1] = t
	}

	jsonData, _ := json.MarshalIndent(tasks, "", " ")
	err := os.WriteFile(FILE_NAME, jsonData, 0644)

	if err != nil {
		panic(err)
	}

	fmt.Println("JSON file written Successfully!")
}

func List() map[int]Task {
	tasks := map[int]Task{}
	// tasks := []Task{}

	data, err := os.ReadFile(FILE_NAME)

	if err == nil && len(data) > 0 {
		_ = json.Unmarshal(data, &tasks)
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

	jsonData, _ := json.MarshalIndent(tasks, "", " ")
	err := os.WriteFile(FILE_NAME, jsonData, 0664)

	if err != nil {
		panic(err)
	}

	deleteMessage := fmt.Sprintf("Delete Task Successfully %s:", id)
	return deleteMessage, nil
}
