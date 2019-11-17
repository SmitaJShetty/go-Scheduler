package model

import (
	"time"

	"github.com/pborman/uuid"

	"github.com/jinzhu/gorm"
)

//Status type for event statuses
type Status string

const (
	//Active represents active status
	Active Status = "active"

	//Inactive represents inactive status
	Inactive Status = "inactive"

	//Stopped represents stopped status
	Stopped Status = "stopped"

	//Pending represents pending status
	Pending Status = "pending"
)

//EventType type for event types -- cron job, one-off, repetitive
type EventType int

const (
	//CronJob that runs every day at a fixed time
	CronJob EventType = 1
	//OneOff job that runs only once at a fixed predefined time
	OneOff EventType = 2
	//Repetitive job that runs every x hours
	Repetitive EventType = 3
)

//Event construct event
type Event struct {
	gorm.Model
	ID        string    `json:"id" gorm:"primary_key"`
	Name      string    `json:"name"`
	Type      EventType `json:"type"`
	Status    Status    `json:"status"`
	Details   []byte    `json:"details"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

//TableName table name
func (e Event) TableName() string {
	return "event"
}

/*
	cronjob: daily, at x hour , y mins
	oneoff : x hour, y mins
	repetitive: every x hour , y mins from first event generation
*/

//NewEvent returns a new event
func NewEvent(eventReq *EventRequest) *Event {
	return &Event{
		ID:        uuid.New(),
		Name:      eventReq.Name,
		Type:      eventReq.Type,
		Status:    eventReq.Status,
		Details:   eventReq.Details,
		CreatedAt: eventReq.CreatedAt,
		UpdatedAt: eventReq.UpdatedAt,
	}
}
