package main

import (
	"crypto/rand"
	"fmt"
	"time"

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

	task2Update := task.Task{
		ID:          rand.Text(),
		Title:       "Update",
		Description: "Update Description",
		Status:      "pending",
		CreatedAt:   time.Now(),
		DueDate:     time.Now().Add(time.Hour * 3),
	}

	// task.Add(task1)
	task.Update("22NCHRYL3J6IKNUGXENMEEN4CL", task2Update)
	message, err := task.Delete("IM4HCFMRG2D7DE3ZD75KIIN6NU")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(message)

	// taskSlice := make([]task.Task, 10)
	// taskSlice[2] = task1

	// taskSlice := []task.Task{}
	// taskList := append(taskSlice, task1)

	// fmt.Println(taskSlice)
	// fmt.Println(len(taskSlice))
	// fmt.Println(taskList)
	// fmt.Println(len(taskList))

	// var taskSlice task.Task = task1 // not a slice btw

	// taskSlice = append(taskSlice, task1)
	// fmt.Println(taskSlice)

}
