package repo

import (
	"scheduler/go-Scheduler/src/scheduler-be/model"

	"github.com/jinzhu/gorm"
)

//EventRepo construct for eventrepo
type EventRepo struct {
	//TODO: logger
	db *gorm.DB
}

//Get get an event based on id
func (e *EventRepo) Get(id string) (*model.Event, error) {
	var event model.Event
	event, err := e.db.Where(model.Event{ID: id}).Find(&event)
	if err != nil {
		return nil, err
	}
	return event, nil
}

//Update updates event
func (e *EventRepo) Update(event *model.Event) (*model.Event, error) {
	database := e.db.Save(event)
	if database.Error != nil {
		return nil, database.Error
	}

	if database.RecordNotFound() {
		return nil, nil
	}
	return event, nil
}

//Delete deletes a event based on id
func (e *EventRepo) Delete(id string) error {
	database := e.db.Delete(&model.Event{ID: id})
	if database.Error != nil {
		return nil, database.Error
	}
}

//Create creates a new event
func (e *EventRepo) Create(event *model.Event) (*model.Event, error) {
	database := e.db.Create(event)
	if database.Error != nil {
		return nil, database.Error
	}

	if database.RecordNotFound() != nil {
		return nil, nil
	}

	return event, nil
}
