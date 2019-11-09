package task_generator

import (
	"fmt"
	"log"
	"scheduler/go-Scheduler/internal/model"
	"scheduler/go-Scheduler/internal/repo"
)

//TaskGenerator tupe for event generator
type TaskGenerator struct {
	//TODO: attach logger
}

//GenerateTasks generate tasks
func (t *TaskGenerator) GenerateTasks(event *model.Event) error {
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
func (t *TaskGenerator) createCronTask(event *model.Event) error {
	if event == nil {
		return fmt.Errorf("event empty")
	}

	fmt.Println("task cron task created")
	return nil
}

//createRepetitiveTask - creates tasks based on event schedule
func (t *TaskGenerator) createRepetitiveTask(event *model.Event) error {
	if event == nil {
		return fmt.Errorf("event empty")
	}

	fmt.Println("event repetitive task")
	return nil
}

//createOneOffTask - creates tasks based on event schedule
func (t *TaskGenerator) createOneOffTask(event *model.Event) error {
	if event == nil {
		return fmt.Errorf("event empty")
	}

	fmt.Println("event one off task created")
	return nil
}

//CreateTasksFromEvents creates tasks for all events
func (t *TaskGenerator) CreateTasksFromEvents() {
	go t.doWork()
}

func (t *TaskGenerator) doWork() {
	//fetch all events
	eventRepo := repo.NewEventRepo()
	events := eventRepo.GetAll()

	//call generate tasks
	for _, event := range events {
		genErr := t.GenerateTasks(&event)
		if genErr != nil {
			//log
			log.Println("error:", genErr.Error())
		}
	}
}

/*
--
cron job: create a scheduled job that runs once every day
an event is created for every cron job
scheduler picks up events and creates tasks
tasks get picked by workers
*/
