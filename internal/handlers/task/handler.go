package task

type TaskHandler struct {
	TaskRepo taskRepository
}

func New(TaskRepo taskRepository) *TaskHandler {
	return &TaskHandler{TaskRepo: TaskRepo}
}
