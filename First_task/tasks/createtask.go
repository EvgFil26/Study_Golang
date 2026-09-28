package tasks

import (
	"time"
)

type Tasks struct {
	Head       string
	Main       string
	Status     bool
	TimeCreate time.Time
	TimeEnd    *time.Time
}

func NewTask(head, main string) *Tasks {
	return &Tasks{
		Head:       head,
		Main:       main,
		Status:     false,
		TimeCreate: time.Now(),
		TimeEnd:    nil,
	}
}

var allTasks []*Tasks

func AddTask(head, main string) *Tasks {
	b := NewTask(head, main)
	allTasks = append(allTasks, b)
	return b
}

func All() []*Tasks {
	return allTasks
}

func Del(head string) []*Tasks {
	for i := 0; i < len(allTasks); i++ {
		if allTasks[i].Head == head {
			allTasks = append(allTasks[:i], allTasks[i+1:]...)
			i--
		}
	}
	return allTasks
}

func Complete(head string) []*Tasks {
	for i := 0; i < len(allTasks); i++ {
		if allTasks[i].Head == head {
			allTasks[i].Status = true
			timeEnd := time.Now()
			allTasks[i].TimeEnd = &timeEnd
		}
	}
	return allTasks
}
