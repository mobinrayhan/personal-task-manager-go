package main

import (
	"mobin.dev/task"
)

func main() {
	// slice version
	// task1 := task.Task{
	// 	ID:          rand.Text(),
	// 	Title:       "Task 1",
	// 	Description: "Task 1 Description",
	// 	Status:      "pending",
	// 	CreatedAt:   time.Now(),
	// 	DueDate:     time.Now().Add(time.Hour * 3),
	// }

	// task2Update := task.Task{
	// 	ID:          rand.Text(),
	// 	Title:       "Update",
	// 	Description: "Update Description",
	// 	Status:      "pending",
	// 	CreatedAt:   time.Now(),
	// 	DueDate:     time.Now().Add(time.Hour * 3),
	// }

	// task.Add(task1)
	// listOfTask := task.List()
	// fmt.Println(listOfTask)
	// task.Update("ZXNGIZCX3CIS7QP2E3K4ZE6L5Q", task2Update)
	task.Delete("57ZUVLDXVCD5T7DJMM5SEHKWOJ")
}
