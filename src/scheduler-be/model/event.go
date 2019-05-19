package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

//EventStatus type for event statuses
type EventStatus string

const (
	//Active event - indicates event has been task generated already
	Active EventStatus = "active"
	//Inactive event - indicates event has not be generated as tasks and needs to be worked on
	Inactive EventStatus = "inactive"
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
	ID        string      `json:"id" gorm:"primary_key"`
	Name      string      `json:"name"`
	Type      EventType   `json:"type"`
	Status    EventStatus `json:"status"`
	Details   []byte      `json:"details"`
	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt time.Time   `json:"updated_at"`
	DeletedAt time.Time   `json:"deleted_at"`
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
