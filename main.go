package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
)

type Task struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

var taskFile string

func LoadTasks() ([]Task, error) {
	data, err := os.ReadFile(taskFile)
	if err != nil {
		if os.IsNotExist(err) {
			return []Task{}, nil // No tasks yet
		}
		return nil, err
	}
	var tasks []Task
	err = json.Unmarshal(data, &tasks)
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func saveTasks(tasks []Task) {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		fmt.Printf("Failed to save task: %v\n", err)
		return
	}

	err = os.MkdirAll("data", os.ModePerm)
	if err != nil {
		fmt.Printf("Failed to create directory: %v\n", err)
		return
	}

	err = os.WriteFile(taskFile, data, 0644)
	if err != nil {
		fmt.Printf("Failed to write file: %v\n", err)
	}
}

func AddTask(title string) error {
	tasks, err := LoadTasks()
	if err != nil {
		return err
	}

	newID := 1
	if len(tasks) > 0 {
		newID = tasks[len(tasks)-1].ID + 1
	}

	newTask := Task{
		ID:        newID,
		Title:     title,
		Completed: false,
	}

	tasks = append(tasks, newTask)
	saveTasks(tasks)
	return nil
}

func MarkTaskDone(id int) error {
	tasks, err := LoadTasks()
	if err != nil {
		return err
	}

	for i := range tasks {
		if tasks[i].ID == id {
			tasks[i].Completed = true
			break
		}
	}

	saveTasks(tasks)
	return nil
}

func DeleteTask(id int) error {
	tasks, err := LoadTasks()
	if err != nil {
		return err
	}

	for i := range tasks {
		if tasks[i].ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			break
		}
	}

	saveTasks(tasks)
	return nil
}

func printTasks(tasks []Task) {
	if len(tasks) == 0 {
		fmt.Println("No tasks found.")
		return
	}
	for _, task := range tasks {
		status := "❌"
		if task.Completed {
			status = "✅"
		}
		fmt.Printf("[%d] %s %s\n", task.ID, status, task.Title)
	}
}

func main() {
	filename := flag.String("file", "data/tasks.json", "Specify the task file")
	addTitle := flag.String("add", "", "Add a new task")
	listFlag := flag.Bool("list", false, "List all tasks")
	doneID := flag.Int("done", 0, "Mark task as done by ID")
	deleteID := flag.Int("delete", 0, "Delete task by ID")

	flag.Parse()
	taskFile = *filename

	if *listFlag {
		tasks, err := LoadTasks()
		if err != nil {
			fmt.Printf("Failed to load tasks: %v\n", err)
			return
		}
		printTasks(tasks)
	}

	if *doneID > 0 {
		if err := MarkTaskDone(*doneID); err != nil {
			fmt.Printf("Error marking task as done: %v\n", err)
		}
	}

	if *deleteID > 0 {
		if err := DeleteTask(*deleteID); err != nil {
			fmt.Printf("Error deleting task: %v\n", err)
		}
	}

	if *addTitle != "" {
		if err := AddTask(*addTitle); err != nil {
			fmt.Printf("Failed to add task: %v\n", err)
		}
	}
}
