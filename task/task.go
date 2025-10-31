package task

import "time"

type Task struct {
	Id          int
	Description string
	Status      string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

const (
	Todo       = "todo"
	InProgress = "in-progress"
	Done       = "done"
)

func NewTask(description string) *Task {
	return &Task{
		Description: description,
		Status:      Todo,
	}
}
