package task

import (
	"encoding/json"
	"fmt"
	"os"
)

type Task struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

var taskFile string

func SetTaskFile(filename string) {
	taskFile = filename
}

func LoadTasks() ([]Task, error) {
	data, err := os.ReadFile(taskFile)
	if err != nil {
		if os.IsNotExist(err) {
			return []Task{}, nil
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
		fmt.Printf("Failed to marshal task: %v\n", err)
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
