package entities

import (
	"fmt"
	"time"
)

type SimpleExecutor struct{}

func New() *SimpleExecutor {
	return &SimpleExecutor{}
}

func (se *SimpleExecutor) Execute(task *Task) {
	fmt.Printf("TASK %d started: %s\n", task.Id, task.Message)
	time.Sleep(time.Duration(task.DelaySec) * time.Second)
	fmt.Printf("TASK %d done\n", task.Id)
}
