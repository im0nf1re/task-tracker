package storage

import (
	"time"

	"github.com/im0nf1re/task-tracker/task"
)

type FileJSONRepository struct {
	driver FileDriver
	tasks  []*task.Task
}

type FileDriver interface {
	ReadTasks(*[]*task.Task) error
	WriteTasks(*[]*task.Task) error
}

func NewFileJSONRepository(fileDriver FileDriver) *FileJSONRepository {
	r := &FileJSONRepository{driver: fileDriver}
	err := r.driver.ReadTasks(&r.tasks)
	if err != nil {
		panic(err)
	}
	return r
}

func (r *FileJSONRepository) Save(t *task.Task) error {
	if t.Id == 0 {
		r.setBaseValuesForCreating(t)
		r.tasks = append(r.tasks, t)
	} else {
		if !r.taskExists(t) {
			r.tasks = append(r.tasks, t)
		}
	}

	err := r.driver.WriteTasks(&r.tasks)
	if err != nil {
		return err
	}

	return nil
}

func (r *FileJSONRepository) Find(id int) (*task.Task, error) {
	for _, el := range r.tasks {
		if el.Id == id {
			return el, nil
		}
	}

	return &task.Task{}, &task.NotFoundError{Id: id}
}

func (r *FileJSONRepository) setBaseValuesForCreating(t *task.Task) {
	var maxId int
	for _, el := range r.tasks {
		if el.Id > maxId {
			maxId = el.Id
		}
	}

	t.Id = maxId + 1

	if t.CreatedAt.IsZero() {
		t.CreatedAt = time.Now()
	}
	if t.UpdatedAt.IsZero() {
		t.UpdatedAt = time.Now()
	}
	if t.Status == "" {
		t.Status = task.Todo
	}
}

func (r *FileJSONRepository) taskExists(t *task.Task) bool {
	for _, el := range r.tasks {
		if el.Id == t.Id {
			return true
		}
	}

	return false
}

func (r *FileJSONRepository) Delete(id int) error {
	for i := range r.tasks {
		if r.tasks[i].Id == id {
			if i == len(r.tasks)-1 {
				r.tasks = r.tasks[:i]
			} else {
				r.tasks = append(r.tasks[:i], r.tasks[i+1:]...)
			}
			err := r.driver.WriteTasks(&r.tasks)
			if err != nil {
				return err
			}
			return nil
		}
	}

	return &task.NotFoundError{Id: id}
}

func (r *FileJSONRepository) List() []*task.Task {
	return r.tasks
}

func (r *FileJSONRepository) ListByStatus(status string) <-chan *task.Task {
	c := make(chan *task.Task)

	go func() {
		defer close(c)
		for _, el := range r.tasks {
			if el.Status == status {
				c <- el
			}
		}
	}()

	return c
}

func (r *FileJSONRepository) ChangeStatus(id int, status string) error {
	t, err := r.Find(id)
	if err != nil {
		return err
	}

	t.Status = status
	err = r.Save(t)
	if err != nil {
		return err
	}

	return nil
}

//func (s *FileJSONRepository) Get() *task.Task {
//	return task.NewTask("test")
//}
//
//func (s *FileJSONRepository) GetAll() []*task.Task {
//	return []*task.Task{}
//}
//
//func (s *FileJSONRepository) GetByStatus(status string) []*task.Task {
//	return []*task.Task{}
//}
//
//func (s *FileJSONRepository) Update(id int, task *task.Task) bool {
//	return true
//}
//
