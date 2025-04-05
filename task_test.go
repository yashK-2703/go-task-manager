package main

import (
	"os"
	"testing"
)

func TestAddAndLoadTask(t *testing.T) {
	taskFile = "test_tasks.json"
	defer os.Remove(taskFile)

	err := AddTask("Test Task")
	if err != nil {
		t.Fatalf("AddTask failed: %v", err)
	}

	tasks, err := LoadTasks()
	if err != nil {
		t.Fatalf("LoadTasks failed: %v", err)
	}

	if len(tasks) != 1 || tasks[0].Title != "Test Task" {
		t.Errorf("Unexpected tasks: %+v", tasks)
	}
}

func TestAddTask(t *testing.T) {
	// Use a test-specific file
	taskFile = "test_tasks.json"
	defer os.Remove(taskFile) // clean up

	err := AddTask("Test something")
	if err != nil {
		t.Fatalf("AddTask failed: %v", err)
	}

	tasks, err := LoadTasks()
	if err != nil {
		t.Fatalf("LoadTasks failed: %v", err)
	}

	if len(tasks) != 1 {
		t.Errorf("Expected 1 task, got %d", len(tasks))
	}
}
