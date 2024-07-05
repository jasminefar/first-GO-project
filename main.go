package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type ToDoList struct {
	tasks []string
}

func (t *ToDoList) AddTask(task string) {
	t.tasks = append(t.tasks, task)
	fmt.Println("Task added:", task)
}

func (t *ToDoList) RemoveTask(index int) {
	if index < 0 || index >= len(t.tasks) {
		fmt.Println("Invalid task number")
		return
	}
	task := t.tasks[index]
	t.tasks = append(t.tasks[:index], t.tasks[index+1:]...)
	fmt.Println("Task removed:", task)
}

func (t *ToDoList) ListTasks() {
	if len(t.tasks) == 0 {
		fmt.Println("No tasks in the list")
		return
	}
	fmt.Println("Tasks:")
	for i, task := range t.tasks {
		fmt.Printf("%d. %s\n", i+1, task)
	}
}

func main() {
	todoList := &ToDoList{}
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\nTo-Do List Menu:")
		fmt.Println("1. Add Task")
		fmt.Println("2. Remove Task")
		fmt.Println("3. List Tasks")
		fmt.Println("4. Exit")
		fmt.Print("Enter your choice: ")

		scanner.Scan()
		choice := scanner.Text()

		switch choice {
		case "1":
			fmt.Print("Enter task: ")
			scanner.Scan()
			task := scanner.Text()
			todoList.AddTask(task)
		case "2":
			fmt.Print("Enter task number to remove: ")
			scanner.Scan()
			indexStr := scanner.Text()
			index, err := strconv.Atoi(indexStr)
			if err != nil {
				fmt.Println("Invalid input")
			} else {
				todoList.RemoveTask(index - 1)
			}
		case "3":
			todoList.ListTasks()
		case "4":
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice, please try again.")
		}
	}
}
