package task

import (
	"errors"
	"fmt"
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
	tasks = append(tasks, t)
	io.WriteToFile(FILE_NAME, tasks)
}

func List() []Task {
	tasks := []Task{}
	io.ReadFromFile(FILE_NAME, &tasks)
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
	io.WriteToFile(FILE_NAME, tasks)
	return id, nil
}

func Delete(id ID) (string, error) {
	tasks := List()

	var itemIndexToRemove int = -1
	for i, v := range tasks {
		fmt.Println(v)
		if v.ID == id {
			itemIndexToRemove = i
			break
		}
	}

	if itemIndexToRemove == -1 {
		formattedMsg := fmt.Sprintf("Task item not found ID : %s", id)
		return "", errors.New(formattedMsg)
	}

	tasks = append(tasks[:itemIndexToRemove], tasks[itemIndexToRemove+1:]...)
	io.WriteToFile(FILE_NAME, tasks)

	deleteMessage := fmt.Sprintf("Delete Task Successfully %s:", id)
	return deleteMessage, nil
}
