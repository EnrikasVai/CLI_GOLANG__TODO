package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"slices"
)

type Task struct {
	Title         string
	DateUntil     string
	DateCompleted string
	Status        bool
}

var tasks []Task

func main() {
	loadFromFile()
	clearTerminal()
	var running bool = true
	greetings()
	for running {
		running = menu()
	}
}
func greetings() {
	fmt.Println("Welcome to your to do list APP")
	fmt.Println("Bellow you can choose from the menu the option you need")
}
func menu() bool {
	var userInput uint8
	fmt.Println("[1] - View tasks that need to be done")
	fmt.Println("[2] - Add a new task")
	fmt.Println("[3] - Delete a task")
	fmt.Println("[4] - Mark task as completed")
	fmt.Println("[5] - Exit the program")
	fmt.Print("Enter the number of task: ")
	fmt.Scan(&userInput)

	switch userInput {
	case 1:
		viewTask()
	case 2:
		addTask()
	case 3:
		deleteTask()
	case 4:
		markTask()
	case 5:
		fmt.Println("Exiting the program, goodbye")
		return false
	default:
		fmt.Println("No such option please select from the menu again")
	}
	return true
}
func viewTask() {
	clearTerminal()
	var taskSubMenu uint8

	if len(tasks) == 0 {
		fmt.Println("You have no task please add task to view them")
		return
	}

	fmt.Println("[1] - View all tasks")
	fmt.Println("[2] - View pending tasks")
	fmt.Println("[3] - View completed tasks")
	fmt.Print("Enter the menu number: ")
	fmt.Scan(&taskSubMenu)

	switch taskSubMenu {
	case 1:
		for i, t := range tasks {
			if t.Status == true {
				fmt.Printf("[%d] - %s Due until - %s, task status = Completed\n", i+1, t.Title, t.DateUntil)
			} else {
				fmt.Printf("[%d] - %s Due until - %s, task status = Pending\n", i+1, t.Title, t.DateUntil)
			}
		}
	case 2:
		if len(tasks) == 0 {
			fmt.Println("You have no task please add task to view them")
			return
		}
		for i, t := range tasks {
			if t.Status == false {
				fmt.Printf("[%d] - %s Due until - %s, task status = Pending\n", i+1, t.Title, t.DateUntil)
			}
		}
	case 3:
		for i, t := range tasks {
			if t.Status == true {
				fmt.Printf("[%d] - %s Due until - %s, task status = Completed\n", i+1, t.Title, t.DateUntil)
			}
		}
	default:
		fmt.Println("No such options please select from the 1,2,3")
	}
}
func addTask() {
	var dateUntil string

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the title of the task:")
	title, _ := reader.ReadString('\n')
	fmt.Print("Enter the date until the task should be finneshed (YYYY-MM-DD): ")
	fmt.Scan(&dateUntil)

	newTask := Task{
		Title:     title,
		DateUntil: dateUntil,
		Status:    false,
	}

	tasks = append(tasks, newTask)
	fmt.Println("Task added successfully")

	saveToFile()
}
func deleteTask() {
	var index int
	if len(tasks) == 0 {
		fmt.Println("You have no task to delete")
		return
	}

	fmt.Println("Select from the task bellow which one you want to delete")
	viewTask()

	fmt.Print("Enter task number you want to delete: ")
	fmt.Scanln(&index)

	if index <= 0 || index > len(tasks) {
		fmt.Println("Invalid task number.")
		return
	}
	tasks = slices.Delete(tasks, index-1, index)
	fmt.Println("Task deleted.")

	saveToFile()
}
func markTask() {
	var index int
	fmt.Println("Select from the task bellow which one you want to mark as complete")

	for i, t := range tasks {
		if t.Status == false {
			fmt.Printf("[%d] - %s Due until - %s, task status = Pending\n", i+1, t.Title, t.DateUntil)
		}
	}

	fmt.Print("Enter task number you want to mark as completed: ")
	fmt.Scanln(&index)

	if index <= 0 || index > len(tasks) {
		fmt.Println("Invalid task number.")
		return
	}
	tasks[index-1].Status = true
	fmt.Println("Task status cahnged.")

	saveToFile()
}
func clearTerminal() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
func saveToFile() {
	file, err := os.Create("tasks.json")
	if err != nil {
		fmt.Println("Error saving file:", err)
		return
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(tasks)
	if err != nil {
		fmt.Println("Error encoding tasks:", err)
	}
}
func loadFromFile() {
	file, err := os.Open("tasks.json")
	if err != nil {
		if os.IsNotExist(err) {
			// File doesn't exist yet, skip loading
			return
		}
		fmt.Println("Error reading file:", err)
		return
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&tasks)
	if err != nil {
		fmt.Println("Error decoding tasks:", err)
	}
}
