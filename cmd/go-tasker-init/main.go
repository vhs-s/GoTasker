package main

import (
	"go-tasker/internal/database"
	"go-tasker/internal/handlers/task"
	"go-tasker/internal/repositories/tasks"
	"net/http"
)

func main() {
	const (
		port string = ":8080"
	)
	// exec := entities.New()
	// tasks := []entities.Task{
	// 	{Id: 1, DelaySec: 5, Message: "Task1"},
	// 	{Id: 2, DelaySec: 5, Message: "Task2"},
	// 	{Id: 3, DelaySec: 5, Message: "Task3"},
	// }
	// cli.Run(tasks, 5, exec)
	db := database.NewCon("database.db")
	repo := tasks.New(db)
	taskHandler := task.New(repo)

	http.HandleFunc("/", taskHandler.IndexHandler())
	http.HandleFunc("/create/", taskHandler.CreateHandler())
	http.HandleFunc("/tasks/", taskHandler.TasksHandler())
	http.ListenAndServe(port, nil)
}
