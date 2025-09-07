package main

import (
	"crypto/rand"
	"fmt"
	"os"
	"strconv"
	"time"

	"mobin.dev/cmd"
	"mobin.dev/task"
)

func main() {

	for {
		chosenOption := cmd.InitiateApp()

		switch chosenOption {
		case 1:
			choseOption := cmd.UserInput("\n 1. Go Default \n 2. Pretty Print \n\nChoose One Option Here ")
			num, _ := strconv.Atoi(choseOption)
			task := task.FetchAllTask(num)
			fmt.Println(task)
		case 2:
			title, description, dueDate := cmd.AllTaskInputFiled()
			task.Add(task.Task{
				ID:          rand.Text(),
				Title:       title,
				Description: description,
				Status:      "pending",
				CreatedAt:   time.Now(),
				DueDate:     dueDate,
			})
		case 3:
			fmt.Println("Update Will Be Added Soon!")
		case 4:
			deletableTaskId := cmd.UserInput("Enter Task Id ")
			task.Delete(deletableTaskId)
		case 5:
			fmt.Println("Thanks For Running That Application. Hope To see You soon!")
			os.Stdin.Close()
		default:
			fmt.Println("Goodbye!!")
			os.Exit(0)
		}

		if !cmd.WannaContinue() {
			fmt.Println("Goodbye!!")
			os.Exit(0)
			break
		}
	}

}
