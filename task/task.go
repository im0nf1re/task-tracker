package task

import "time"

type Task struct {
	Id          int
	Description string
	Status      string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func NewTask(description string) *Task {
	return &Task{
		Description: description,
		Status:      TODO,
	}
}
