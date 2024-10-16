package main

import (
	"fmt"
)

func main() {
	var todoList []string
	var choice int

	for {
		fmt.Println("\n-- To-Do List Menu --")
		fmt.Println("1. Add Task")
		fmt.Println("2. View Tasks")
		fmt.Println("3. Delete Task")
		fmt.Println("4. Mark Task as Completed")
		fmt.Println("5. Exit")
		fmt.Print("Enter your choice: ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			addTask(&todoList)
		case 2:
			viewTasks(todoList)
		case 3:
			deleteTask(&todoList)
		case 4:
			markCompleted(&todoList)
		case 5:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}

// Function to add a task
func addTask(todoList *[]string) {
	var task string
	fmt.Print("Enter the task: ")
	fmt.Scanln(&task)
	*todoList = append(*todoList, task)
	fmt.Println("Task added successfully!")
}

// Function to view all tasks
func viewTasks(todoList []string) {
	if len(todoList) == 0 {
		fmt.Println("No tasks available.")
		return
	}
	fmt.Println("\nYour tasks:")
	for i, task := range todoList {
		fmt.Printf("%d. %s\n", i+1, task)
	}
}

// Function to delete a task by index
func deleteTask(todoList *[]string) {
	viewTasks(*todoList)
	if len(*todoList) == 0 {
		return
	}
	var index int
	fmt.Print("Enter the number of the task to delete: ")
	fmt.Scanln(&index)

	if index > 0 && index <= len(*todoList) {
		*todoList = append((*todoList)[:index-1], (*todoList)[index:]...)
		fmt.Println("Task deleted!")
	} else {
		fmt.Println("Invalid task number.")
	}
}

// Function to mark a task as completed
func markCompleted(todoList *[]string) {
	viewTasks(*todoList)
	if len(*todoList) == 0 {
		return
	}
	var index int
	fmt.Print("Enter the number of the task to mark as completed: ")
	fmt.Scanln(&index)

	if index > 0 && index <= len(*todoList) {
		(*todoList)[index-1] = "[X] " + (*todoList)[index-1]
		fmt.Println("Task marked as completed!")
	} else {
		fmt.Println("Invalid task number.")
	}
}
