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
	createTaskChan := make(chan string)
	getTaskChan := make(chan interface{})

	for {
		chosenOption := cmd.InitiateApp()

		switch chosenOption {
		case 1:
			choseOption := cmd.UserInput("\n 1. Go Default \n 2. Pretty Print \n\nChoose One Option Here ")
			num, _ := strconv.Atoi(choseOption)
			go task.FetchAllTask(num, getTaskChan)
			fmt.Println(<-getTaskChan)
		case 2:
			title, description, dueDate := cmd.AllTaskInputFiled()
			go task.Add(task.Task{
				ID:          rand.Text(),
				Title:       title,
				Description: description,
				Status:      "pending",
				CreatedAt:   time.Now(),
				DueDate:     dueDate,
			}, createTaskChan)
			go task.FetchAllTask(2, getTaskChan)
			fmt.Println(<-getTaskChan)
			fmt.Println(<-createTaskChan)
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

// package main

// import (
// 	"fmt"
// 	"time"
// )

// func greet(phrase string, doneChan chan string) {
// 	fmt.Println("Hello!", phrase)
// 	doneChan <- "Greet Done!"
// }

// func slowGreet(phrase string, doneChan chan string) {
// 	time.Sleep(3 * time.Second) // simulate a slow, long-taking task
// 	fmt.Println("Hello!", phrase)
// 	doneChan <- "Slow Greet Done!"
// 	close(doneChan) // close the channel to signal completion
// }

// func main() {
// 	// dones := make([]chan string, 4)

// 	done := make(chan string)
// 	// dones[0] = make(chan string)
// 	go greet("Nice to meet you!", done)
// 	// dones[1] = make(chan string)
// 	go greet("How are you?", done)
// 	// dones[2] = make(chan string)
// 	go slowGreet("How ... are ... you ...?", done)
// 	// dones[3] = make(chan string)
// 	go greet("I hope you're liking the course!", done)
// 	// <-dones[0]
// 	// <-dones[1]
// 	// <-dones[2]
// 	// <-dones[3]

// 	for range done {
// 		<-done
// 	}
// }
