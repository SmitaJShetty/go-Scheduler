package task_generator

import (
	"fmt"
	"scheduler/go-Scheduler/src/scheduler-be/model"
)

//EventGenerator tupe for event generator
type TaskGenerator struct {
	//TODO: attach logger
}

//GenerateTasks generate events
func (e *EventGenerator) GenerateTasks(event *model.Event) error {
	//Create events based on request
	if event == nil {
		return fmt.Errorf("events are empty")
	}

	if event.Status == model.Active {
		return fmt.Errorf("event is active")
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
func (e *EventGenerator) createCronTask(event *model.Event) error {
	if event == nil {
		return fmt.Errorf("event empty")
	} 

	taskRepo:= repo.NewTaskRepo()
	taskRepo.
	return nil
}

//createRepetitiveTask - creates tasks based on event schedule
func (e *EventGenerator) createRepetitiveTask(event *model.Event) error {
	if event == nil {
		return fmt.Errorf("event empty")
	}

	return nil
}

//createOneOffTask - creates tasks based on event schedule
func (e *EventGenerator) createOneOffTask(event *model.Event) error {
	if event == nil {
		return fmt.Errorf("event empty")
	}

	return nil
}

/*
-- cron job: create a scheduled job that runs once every day

an event is created for every cron job

scheduler picks up events and creates tasks

tasks get picked by workers
*/
