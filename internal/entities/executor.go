package entities

import (
	"fmt"
	"time"
)

type SimpleExecutor struct{}

func NewSimpleExec() *SimpleExecutor {
	return &SimpleExecutor{}
}

func (se *SimpleExecutor) Execute(task *Task) {
	fmt.Printf("TASK %s started: %s\n", task.Id, task.Message)
	time.Sleep(time.Duration(task.DelaySec) * time.Second)
	fmt.Printf("TASK %s done\n", task.Id)
}
