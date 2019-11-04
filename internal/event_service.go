package internal

import "scheduler/go-Scheduler/internal/model"

type EventService struct {
}

//NewEventService returns a NewEventService
func NewEventService() *EventService {
	return &EventService{}
}

// CreateEvent creates an event
func (es *EventService) CreateEvent(eventReq *model.EventRequest) error {
	//create event
	event := model.NewEvent(eventReq)
	eventRepo := repo.NewEventRepo()
	createEventErr := eventRepo.CreateEvent(event)
	if createEventErr != nil {
		return createEventErr
	}

	return nil
}
