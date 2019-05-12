package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

//Event construct event
type Event struct {
	gorm.Model
	ID        string    `json:"id" gorm:"primary_key"`
	Name      string    `json:"name"`
	Type      string    `json:"type"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

//TableName table name
func (e Event) TableName() string {
	return "event"
}
