package internal

import (
	"fmt"

	"github.com/smitajshetty/go-scheduler/internal/model"
	"github.com/smitajshetty/go-scheduler/internal/repo"
)

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
	newEvent, createEventErr := eventRepo.Create(event)
	if createEventErr != nil {
		return createEventErr
	}

	if newEvent == nil {
		return fmt.Errorf("no event was created")
	}
	return nil
}
