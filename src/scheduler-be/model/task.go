package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

//Task construct fr tasks
type Task struct {
	gorm.Model
	ID        string    `json:"id" gorm:"primary_key"`
	Name      string    `json:"name"`
	Type      string    `json:"type"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

//TableName name of task
func (t Task) TableName() string {
	return "task"
}
