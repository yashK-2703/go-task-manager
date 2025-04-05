package main

import (
	"flag"
	"fmt"

	"github.com/yashK-2703/go-task-manager/task"
)

func main() {
	fmt.Println("Task manager starting...")

	addTitle := flag.String("add", "", "Add a new task")
	listFlag := flag.Bool("list", false, "List all tasks")
	doneID := flag.Int("done", 0, "Mark task as done by ID")
	deleteID := flag.Int("delete", 0, "Delete task by ID")
	filename := flag.String("file", "data/tasks.json", "Specify the task file")

	flag.Parse()

	task.SetTaskFile(*filename)

	if *listFlag {
		tasks, err := task.LoadTasks()
		if err != nil {
			fmt.Printf("Failed to load tasks: %v\n", err)
			return
		}
		if len(tasks) == 0 {
			fmt.Println("No tasks found.")
			return
		}

		for _, t := range tasks {
			status := "❌"
			if t.Completed {
				status = "✅"
			}
			fmt.Printf("[%d] %s %s\n", t.ID, status, t.Title)
		}
	}

	if *doneID > 0 {
		task.MarkTaskDone(*doneID)
	}

	if *deleteID > 0 {
		task.DeleteTask(*deleteID)
	}

	if *addTitle != "" {
		err := task.AddTask(*addTitle)
		if err != nil {
			fmt.Printf("Failed to add task %v\n", err)
		}
	}
}
