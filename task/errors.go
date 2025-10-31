package task

import "fmt"

type NotFoundError struct {
	Id int
}

func (e *NotFoundError) Error() string {
	return fmt.Sprintf("Task with id %d not found:", e.Id)
}
