package cli

import (
	"go-tasker/internal/entities"
	"sync"
)

func Run(tasks []entities.Task, workersCount int, executor entities.TaskExecutor) {
	taskchan := make(chan entities.Task)
	var wg sync.WaitGroup
	for i := 0; i < workersCount; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for diegotask := range taskchan {
				executor.Execute(&diegotask)
			}
		}()
	}
	go func() {
		for _, task := range tasks {
			taskchan <- task
		}
		close(taskchan)
	}()
	wg.Wait()
}
