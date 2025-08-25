package entities

import (
	"log"
	"strconv"

	"github.com/google/uuid"
)

const (
	StatusRunning   = "RUNNING"
	StatusNotRun    = "NOTRUN"
	StatusCompleted = "Completed"
)

type Task struct {
	Id       string
	Name     string
	DelaySec int
	Message  string
	Status   string
}

func New(name string, delaySec int, message string) *Task {
	uuid := uuid.New().String()
	return &Task{
		Id:       uuid,
		Name:     name,
		DelaySec: delaySec,
		Message:  message,
		Status:   StatusNotRun,
	}
}

func DelayConv(delayStr string) int {
	convertedDelay, err := strconv.Atoi(delayStr)
	if err != nil {
		log.Println("Error to convert delay")
	}
	return convertedDelay
}
