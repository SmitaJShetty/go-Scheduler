package model

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/pborman/uuid"
)

//Task construct fr tasks
type Task struct {
	gorm.Model
	ID        string    `json:"id" gorm:"primary_key"`
	Name      string    `json:"name"`
	EventID   string    `json:"event_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

//TableName name of task
func (t Task) TableName() string {
	return "task"
}

//NewTask returns a new task
func NewTask(name string, eventID string) *Task {
	return &Task{
		Name:    name,
		EventID: eventID,
		ID:      uuid.NewRandom(),
	}
}

/*

task generator -- creates tasks from events

talk to db
plug in kafka -- put tasks from events
deploy on aws
get pipelines done
GET UI DONE

workers -- another service, picks up tasks, works
THings to get done by Monday evening:




*/
