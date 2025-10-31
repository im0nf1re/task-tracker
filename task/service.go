package task

import "fmt"

type Storage interface {
	Save(t *Task) error
	Find(id int) (*Task, error)
	Delete(id int) error
	List() []*Task
	ListByStatus(status string) <-chan *Task
	ChangeStatus(id int, status string) error
}

type Service struct {
	Storage Storage
}

func NewService(storage Storage) *Service {
	return &Service{Storage: storage}
}

func (s *Service) Add(description string) (*Task, error) {
	t := NewTask(description)

	err := s.Storage.Save(t)
	if err != nil {
		return &Task{}, err
	}

	return t, nil
}

func (s *Service) Update(id int, description string) (*Task, error) {
	t, err := s.Storage.Find(id)
	if err != nil {
		return &Task{}, err
	}
	t.Description = description

	err = s.Storage.Save(t)
	if err != nil {
		return &Task{}, err
	}

	return t, nil
}

func (s *Service) Delete(id int) error {
	return s.Storage.Delete(id)
}

func (s *Service) List() {
	for _, t := range s.Storage.List() {
		fmt.Print("ID: ", t.Id, " ", t.Status, " \"", t.Description, "\"\n")
	}
}

func (s *Service) ListByStatus(status string) {
	for t := range s.Storage.ListByStatus(status) {
		fmt.Print("ID: ", t.Id, " ", t.Status, " \"", t.Description, "\"\n")
	}
}

func (s *Service) MarkInProgress(id int) error {
	return s.Storage.ChangeStatus(id, InProgress)

}

func (s *Service) MarkDone(id int) error {
	return s.Storage.ChangeStatus(id, Done)
}
