package tasks

const (
	QueryInsertTask   = "INSERT INTO tasks (Id, TaskName, TaskDelay, TaskMessage, Stat) values (?,?,?,?,?)"
	QueryGetTasks     = "SELECT * FROM tasks"
	QueryCountAllRows = "SELECT COUNT(*) FROM tasks"
)
