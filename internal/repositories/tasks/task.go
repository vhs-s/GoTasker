package tasks

import (
	"database/sql"
	"go-tasker/internal/entities"
	"log"
)

type Task struct {
	Id       string
	Name     string
	DelaySec int
	Message  string
	Status   string
}

type repository struct {
	db *sql.DB
}

func New(db *sql.DB) *repository {
	return &repository{db: db}
}

func (r *repository) Create(task *entities.Task) error {
	query := QueryInsertTask
	_, err := r.db.Exec(query, task.Id, task.Name, task.DelaySec, task.Message, task.Status)
	if err != nil {
		log.Println("Error creating record in database:", err)
		return err
	}
	return nil
}
func (r *repository) Get() []Task {
	var count int
	query := QueryGetTasks
	query2 := QueryCountAllRows
	row := r.db.QueryRow(query2)
	row.Scan(&count)

	tasks := make([]Task, 0, count)

	rows, err := r.db.Query(query)
	if err != nil {
		log.Println(err)
	}
	for rows.Next() {
		var id, name, message, status string
		var delay int

		err := rows.Scan(&id, &name, &delay, &message, &status)
		if err != nil {
			log.Println(err)
		}
		task := Task{Id: id, Name: name, DelaySec: delay, Message: message, Status: status}
		tasks = append(tasks, task)
	}
	defer rows.Close()
	return tasks
}
