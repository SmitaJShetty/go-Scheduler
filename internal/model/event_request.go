package model

import (
	"github.com/jinzhu/gorm"
)

//EventRequest construct event
type EventRequest struct {
	gorm.Model
	Name    string       `json:"name"`
	Type    EventType    `json:"type"`
	Status  Status       `json:"status"`
	Details EventDetails `json:"details"`
}

//EventDetails construct event
type EventDetails struct {
	Location string `json:"location"`
}
