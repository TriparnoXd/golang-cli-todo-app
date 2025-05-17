package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

type Task struct {
	Task string `json:"task"`
	Done bool   `json:"done"`
}

const fileName = "todo.json"

// Load tasks from file
func loadTasks() ([]Task, error) {
	var tasks []Task

	data, err := os.ReadFile(fileName)
	if err != nil {
		// If file doesn't exist, return empty list
		if os.IsNotExist(err) {
			return tasks, nil
		}
		return nil, err
	}

	err = json.Unmarshal(data, &tasks)
	return tasks, err
}

// Save tasks to file
func saveTasks(tasks []Task) error {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(fileName, data, 0644)
}

// Add a task
func addTask(taskDesc string) {
	tasks, _ := loadTasks()
	tasks = append(tasks, Task{Task: taskDesc, Done: false})
	saveTasks(tasks)
	fmt.Println("Task added.")
}

// List all tasks
func listTasks() {
	tasks, _ := loadTasks()
	for i, t := range tasks {
		status := " "
		if t.Done {
			status = "✓"
		}
		fmt.Printf("%d. [%s] %s\n", i+1, status, t.Task)
	}
}

// Mark task as done
func completeTask(index int) {
	tasks, _ := loadTasks()
	if index < 1 || index > len(tasks) {
		fmt.Println("Invalid task number.")
		return
	}
	tasks[index-1].Done = true
	saveTasks(tasks)
	fmt.Println("Task marked as done.")
}

func deleteTask(index int) {
	tasks, _ := loadTasks()
	if index < 1 || index > len(tasks) {
		fmt.Println("Invalid task number.")
		return
	}
	tasks = append(tasks[:index-1], tasks[index:]...)
	saveTasks(tasks)
	fmt.Println("Task deleted.")
}

// CLI entry point
func main() {
	args := os.Args

	if len(args) < 2 {
		fmt.Println("Usage: todo [add|list|done|delete] [task]")
		return
	}

	command := args[1]

	switch command {
	case "add":
		if len(args) < 3 {
			fmt.Println("Provide a task to add.")
			return
		}
		addTask(args[2])
	case "list":
		listTasks()
	case "done":
		if len(args) < 3 {
			fmt.Println("Provide the task number to mark as done.")
			return
		}
		i, _ := strconv.Atoi(args[2])
		completeTask(i)
	case "delete":
		if len(args) < 3 {
			fmt.Println("Provide the task number to delete.")
			return
		}
		i, _ := strconv.Atoi(args[2])
		deleteTask(i)
	default:
		fmt.Println("Unknown command:", command)
	}
}
