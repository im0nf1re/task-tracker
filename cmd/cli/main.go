package main

import (
	"os"

	"github.com/im0nf1re/task-tracker/storage"
	"github.com/im0nf1re/task-tracker/task"
)

func main() {
	args := os.Args[1:]
	fileJSONStorage := storage.NewFileJSONStorage()
	handler := task.NewHandler(fileJSONStorage)

	switch args[0] {
	case "add":
		handler.Add(args[1])
		//case "update":
		//
		//case "delete":
		//
		//case "mark-in-progress":
		//
		//case "mark-done":
		//
		//case "list":
	}
}
