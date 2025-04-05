package task

import (
	"os"
	"testing"
)

func TestAddTask(t *testing.T) {
	SetTaskFile("test_tasks.json")
	os.Remove("test_tasks.json") // clean start

	err := AddTask("Test Task")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	tasks, _ := LoadTasks()
	if len(tasks) != 1 || tasks[0].Title != "Test Task" {
		t.Fatalf("Expected task to be added")
	}
}
