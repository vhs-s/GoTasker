package main

import (
	"go-tasker/internal/cli"
	"go-tasker/internal/entities"
)

func main() {
	exec := entities.New()
	tasks := []entities.Task{
		{Id: 1, DelaySec: 5, Message: "Task1"},
		{Id: 2, DelaySec: 5, Message: "Task2"},
		{Id: 3, DelaySec: 5, Message: "Task3"},
	}
	cli.Run(tasks, 5, exec)
}
