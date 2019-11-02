package internal

import (
	"fmt"
	"scheduler/go-Scheduler/internal/model"
)

//NewEventGenerator returns a Taskgenerator
func NewEventGenerator() *TaskGenerator {
	return &TaskGenerator{}
}

//TaskGenerator tupe for event generator
type TaskGenerator struct {
	//TODO: attach logger
}

//CreateTasks generate tasks
func (eg *TaskGenerator) CreateTasks(event *model.Event) error {
	//Create events based on request
	if event == nil {
		return fmt.Errorf("events are empty")
	}

	switch event.Type {
	case model.CronJob:
		cronJobErr := e.createCronTask(event)
		if cronJobErr != nil {
			return cronJobErr
		}
	case model.Repetitive:
		repetitiveJobErr := e.createRepetitiveTask(event)
		if repetitiveJobErr != nil {
			return repetitiveJobErr
		}
		break
	case model.OneOff:
		oneOffJobErr := e.createOneOffTask(event)
		if oneOffJobErr != nil {
			return oneOffJobErr
		}
	default:
		return fmt.Errorf("invalid job type")
	}
	return nil
}

//createCronTask - creates tasks based on event schedule
func (eg *TaskGenerator) createCronTask(event *model.Event) error {
	if event == nil {
		return fmt.Errorf("event empty")
	}

	taskRepo := repo.NewTaskRepo()
	return nil
}

//createRepetitiveTask - creates tasks based on event schedule
func (eg *TaskGenerator) createRepetitiveTask(event *model.Event) error {
	if event == nil {
		return fmt.Errorf("event empty")
	}

	return nil
}

//createOneOffTask - creates tasks based on event schedule
func (eg *TaskGenerator) createOneOffTask(event *model.Event) error {
	if event == nil {
		return fmt.Errorf("event empty")
	}

	return nil
}

//StartTaskGenerator start task generator routine
func StartTaskGenerator() error {
	eventGen := NewEventGenerator()
	go func() {
		eventGen.CreateAllTasks()
	}()

	return nil
}

// CreateAllTasks creates all tasks
func (eg *TaskGenerator) CreateAllTasks() {
	//fetch all types of= tasks
	eventRepo := repo.NewEventRepo()

	var eventList []model.Event
	eventList = eventRepo.GetAllEvents()

	//run through create task function
	for i, ev := range eventList {

	}
}
