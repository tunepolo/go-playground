package main

import "fmt"

// ID of Task
type ID int

// Priority of Task
type Priority int

// Task Struct definition
type Task struct {
	id       ID
	priority Priority
	detail   string
	done     bool
}

// NewTask is a pseudo Constructor
func NewTask(id ID, detail string) *Task {
	task := &Task{
		id:       id,
		priority: 0,
		detail:   detail,
		done:     false,
	}
	return task
}

// Receiver
func (task Task) String() string {
	str := fmt.Sprintf("%d) %s", task.id, task.detail)
	return str
}

func main() {
	task := NewTask(1, "Sample Task")
	fmt.Println(task)
}
