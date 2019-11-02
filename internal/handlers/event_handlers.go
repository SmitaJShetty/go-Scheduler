package handlers

import (
	"encoding/json"
	"net/http"
	"scheduler/go-Scheduler/internal/model"
	"scheduler/go-Scheduler/pkg/common"
)

// CreateEventHandler handler for events
func CreateEventHandler(w http.ResponseWriter, r *http.Request) {
	var eventRequest model.Event
	err := json.NewDecoder(r.Body).Decode(&eventRequest)
	if err != nil {
		common.SendErrorResponse(w, r, common.NewAppError(err.Error(), http.StatusBadRequest))
		return
	}

	createEventErr := service.NewEventService().CreateEvent(&eventRequest)
	if createEventErr != nil {
		common.SendErrorResponse(w, r, common.NewAppError(createEventErr.Error(), http.StatusInternalServerError))
		return
	}
	common.SendResult(w, r, nil)
	return
}
