package storage

import (
	"encoding/json"
	"os"
	"time"

	"github.com/im0nf1re/task-tracker/task"
)

type FileJSONStorage struct {
	tasks []*task.Task
}

func NewFileJSONStorage() *FileJSONStorage {
	var tasks []*task.Task

	file, err := os.ReadFile("C:\\Users\\imonfire\\projects\\go\\task-tracker\\var\\storage.txt")
	if err != nil {
		return &FileJSONStorage{tasks}
	}

	err = json.Unmarshal(file, &tasks)
	if err != nil {
		panic(err)
	}

	return &FileJSONStorage{tasks}
}

func (s *FileJSONStorage) Create(t *task.Task) error {
	var maxId int
	for _, el := range s.tasks {
		if el.Id > maxId {
			maxId = el.Id
		}
	}

	t.Id = maxId + 1
	t.CreatedAt = time.Now()
	t.UpdatedAt = time.Now()
	t.Status = task.TODO

	s.tasks = append(s.tasks, t)

	tasksByte, err := json.Marshal(s.tasks)
	if err != nil {
		return err
	}
	err = os.WriteFile("C:\\Users\\imonfire\\projects\\go\\task-tracker\\var\\storage.txt", tasksByte, 0644)
	if err != nil {
		return err
	}

	return nil
}

//func (s *FileJSONStorage) Get() *task.Task {
//	return task.NewTask("test")
//}
//
//func (s *FileJSONStorage) GetAll() []*task.Task {
//	return []*task.Task{}
//}
//
//func (s *FileJSONStorage) GetByStatus(status string) []*task.Task {
//	return []*task.Task{}
//}
//
//func (s *FileJSONStorage) Update(id int, task *task.Task) bool {
//	return true
//}
//
//func (s *FileJSONStorage) Delete(id int) bool {
//	return true
//}
