package task

import (
	"go-tasker/internal/entities"
	"go-tasker/internal/repositories/tasks"
)

type taskRepository interface {
	Create(task *entities.Task) error
	Get() []tasks.Task
}
