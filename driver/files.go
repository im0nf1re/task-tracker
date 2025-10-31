package driver

import (
	"encoding/json"
	"os"

	"github.com/im0nf1re/task-tracker/task"
)

type Files struct {
	filePath string
}

func NewFiles(filePath string) *Files {
	return &Files{filePath: filePath}
}

func (f *Files) ReadTasks(taskSlicePointer *[]*task.Task) error {
	file, err := os.ReadFile(f.filePath)
	if err != nil {
		return nil
	}

	if len(file) == 0 {
		return nil
	}

	err = json.Unmarshal(file, taskSlicePointer)
	if err != nil {
		return err
	}

	return nil
}

func (f *Files) WriteTasks(taskSlicePointer *[]*task.Task) error {
	tasksByte, err := json.Marshal(taskSlicePointer)
	if err != nil {
		return err
	}
	err = os.WriteFile(f.filePath, tasksByte, 0644)
	if err != nil {
		return err
	}

	return nil
}
