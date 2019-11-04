package model

import (
	"github.com/jinzhu/gorm"
)

//EventRequest construct event
type EventRequest struct {
	gorm.Model
	Name    string      `json:"name"`
	Type    EventType   `json:"type"`
	Status  EventStatus `json:"status"`
	Details []byte      `json:"details"`
}
