package repo

import "scheduler/go-Scheduler/src/scheduler-be/model"

//EventRepo construct for eventrepo
type EventRepo struct {
}

//Get get an event based on id
func (e *EventRepo) Get(id string) (*model.Event, error) {
	return nil, nil
}

//Update updates event
func (e *EventRepo) Update(event *model.Event) (*model.Event, error) {
	return nil, nil
}

//Delete deletes a event based on id
func (e *EventRepo) Delete(id string) (*model.Event, error) {
	return nil, nil
}

//Create creates a new event
func (e *EventRepo) Create(event *model.Event) error {
	return nil
}
