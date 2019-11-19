package event_generator

import (
	"fmt"

	"github.com/smitajshetty/go-scheduler/internal/model"
)

//EventGenerator construct for  event generator
type EventGenerator struct {
	//TODO: logger
}

//GenerateEvents generate events
func (e *EventGenerator) GenerateEvents(event *model.Event) error {
	if event == nil {
		return fmt.Errorf("event was empty")
	}
	return nil
}
