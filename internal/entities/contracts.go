package entities

type TaskExecutor interface {
	Execute(task *Task)
}
