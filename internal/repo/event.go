package repo

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/smitajshetty/go-scheduler/internal/model"
)

const server string = "0.0.0.0:5432"
const connString string = "host=0.0.0.0 port=5432 user=smita dbname=scheduler password=pwd123"

//NewEventRepo starts a new event repo
func NewEventRepo() *EventRepo {
	newDB, dbErr := gorm.Open(server, connString)
	if dbErr != nil {
		//log
		fmt.Println(dbErr.Error())
		return nil
	}

	return &EventRepo{
		db: newDB,
	}
}

//EventRepo construct for eventrepo
type EventRepo struct {
	//TODO: logger
	db *gorm.DB
}

//GetAll get all events
func (e *EventRepo) GetAll() ([]*model.Event, error) {
	var events []*model.Event
	db := e.db.Where(model.Event{}).Find(events)
	if db.Error != nil {
		return nil, db.Error
	}
	return events, nil
}

//Get get an event based on id
func (e *EventRepo) Get(id string) (*model.Event, error) {
	var event model.Event
	db := e.db.Where(model.Event{ID: id}).Find(&event)
	if db.Error != nil {
		return nil, db.Error
	}
	return &event, nil
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
		return database.Error
	}

	event, getErr := e.Get(id)
	if getErr != nil {
		return getErr
	}

	if event == nil {
		return fmt.Errorf("event is not found for id:%s", id)
	}

	database = e.db.Delete(event)
	if database.Error != nil {
		return database.Error
	}

	return nil
}

//Create creates a new event
func (e *EventRepo) Create(event *model.Event) (*model.Event, error) {
	database := e.db.Create(event)
	if database.Error != nil {
		return nil, database.Error
	}

	if database.RecordNotFound() {
		return nil, nil
	}

	return event, nil
}
