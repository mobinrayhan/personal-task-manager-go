package main

import (
	"crypto/rand"
	"fmt"
	"time"

	"mobin.dev/task"
)

func main() {
	// slice version
	task1 := task.Task{
		ID:          rand.Text(),
		Title:       "Task 1",
		Description: "Task 1 Description",
		Status:      "pending",
		CreatedAt:   time.Now(),
		DueDate:     time.Now().Add(time.Hour * 3),
	}

	task2Update := task.Task{
		ID:          rand.Text(),
		Title:       "Update",
		Description: "Update Description",
		Status:      "pending",
		CreatedAt:   time.Now(),
		DueDate:     time.Now().Add(time.Hour * 3),
	}

	task.Add(task1)
	listOfTask := task.List()
	fmt.Println(listOfTask)
	task.Update("6H4DHPJYUK2HWTGOQSRIWLIRNL", task2Update)
	task.Delete("CGPE2PZ2JODYZ7547AWVYCELFL")
}
