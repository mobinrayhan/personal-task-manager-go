package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"mobin.dev/util"
)

func UserInput(labeText string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(labeText + ": ")
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Failed To Read User Input : ", err)
		return ""
	}
	cleanedInput := strings.TrimSuffix(input, "\n")
	cleanedInput = strings.TrimSuffix(input, "\r")
	cleanedInput = strings.TrimSpace(input)
	return cleanedInput
}

func WannaContinue() bool {
	choseOption := UserInput("\n 1. Continue \n 2. Cancel \n\nChoose One Continue Option Here ")
	intN, err := strconv.Atoi(choseOption)

	if err != nil || intN > 2 || intN < 0 {
		fmt.Println("Not A Valid Number , Number Should Under 5 and Upper 1")
		os.Exit(0)
	}

	return intN == 1
}

func InitiateApp() int {
	choseOption := UserInput("\n 1. Fetch All Task \n 2. Add New Task \n 3. Update Task \n 4. Delete Task \n 5. Cancel \n\nChoose One Option Here ")

	intN, err := strconv.Atoi(choseOption)

	if err != nil || intN > 5 || intN < 0 {
		fmt.Println("Not A Valid Number , Number Should Under 5 and Upper 1")
		os.Exit(0)
	}

	return intN
}

func AllTaskInputFiled() (string, string, time.Time) {
	title := UserInput("Enter Task Name")
	description := UserInput("Enter Task Description")
	dueDate := UserInput("Enter Due Date dd/mm/yyyy")

	t, isOkay := util.IsValidDate(dueDate)

	if !isOkay {
		panic("input is not valid date format dd/mm/yyyy")
	}

	return title, description, t
}
