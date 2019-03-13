package main

import "fmt"

// ID of Task
type ID int

// Priority of Task
type Priority int

// User name
type User struct {
	FirstName string
	LastName  string
}

// FullName return user full name
func (u *User) FullName() string {
	return fmt.Sprintf("%s %s", u.FirstName, u.LastName)
}

// NewUser return new user object
func NewUser(firstName string, lastName string) *User {
	return &User{
		FirstName: firstName,
		LastName:  lastName,
	}
}

// Task Struct definition
type Task struct {
	id       ID
	priority Priority
	detail   string
	done     bool
	*User
}

// NewTask is a pseudo Constructor
func NewTask(id ID, detail string, firstName string, lastName string) *Task {
	task := &Task{
		id:       id,
		priority: 0,
		detail:   detail,
		done:     false,
		User:     NewUser(firstName, lastName),
	}
	return task
}

// Receiver
func (task Task) String() string {
	str := fmt.Sprintf("(%d) %s -> %s", task.id, task.detail, task.FullName())
	return str
}

func main() {
	task := NewTask(1, "Sample Task", "Yuichi", "Tsunematsu")
	fmt.Println(task)
}
