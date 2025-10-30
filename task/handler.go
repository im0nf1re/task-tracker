package task

type TaskStorage interface {
	Create(t *Task) error
	//Get() (*task.Task, error)
	//GetAll() ([]*task.Task, error)
	//GetByStatus(status string) ([]*task.Task, error)
	//Update(id int, task *task.Task) (int, error)
	//Delete(id int) (bool, error)
}

type Handler struct {
	Storage TaskStorage
}

func NewHandler(storage TaskStorage) *Handler {
	return &Handler{Storage: storage}
}

func (h *Handler) Add(description string) *Task {
	t := NewTask(description)

	err := h.Storage.Create(t)
	if err != nil {
		panic(err)
	}

	return t
}
