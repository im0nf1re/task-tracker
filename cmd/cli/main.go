package main

import (
	"os"
	"strconv"

	"github.com/im0nf1re/task-tracker/driver"
	"github.com/im0nf1re/task-tracker/storage"
	"github.com/im0nf1re/task-tracker/task"
)

func main() {
	args := os.Args[1:]
	fileDriver := driver.NewFiles("C:\\Users\\imonfire\\projects\\go\\task-tracker\\var\\storage.txt")
	fileJSONRepository := storage.NewFileJSONRepository(fileDriver)
	taskService := task.NewService(fileJSONRepository)

	switch args[0] {
	case "add":
		_, err := taskService.Add(args[1])
		if err != nil {
			panic(err)
		}
		taskService.List()
	case "update":
		arg1, err := strconv.Atoi(args[1])
		if err != nil {
			panic(err)
		}
		_, err = taskService.Update(arg1, args[2])
		if err != nil {
			panic(err)
		}
		taskService.List()
	case "delete":
		arg1, err := strconv.Atoi(args[1])
		if err != nil {
			panic(err)
		}
		err = taskService.Delete(arg1)
		if err != nil {
			panic(err)
		}
		taskService.List()
	case "list":
		if len(args) > 1 {
			taskService.ListByStatus(args[1])
		} else {
			taskService.List()
		}
	case "mark-in-progress":
		arg1, err := strconv.Atoi(args[1])
		if err != nil {
			panic(err)
		}
		err = taskService.MarkInProgress(arg1)
		if err != nil {
			return
		}
		taskService.List()
	case "mark-done":
		arg1, err := strconv.Atoi(args[1])
		if err != nil {
			panic(err)
		}
		err = taskService.MarkDone(arg1)
		if err != nil {
			return
		}
		taskService.List()
	}
}
