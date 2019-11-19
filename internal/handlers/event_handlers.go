package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/smitajshetty/go-scheduler/internal"
	"github.com/smitajshetty/go-scheduler/internal/model"
	"github.com/smitajshetty/go-scheduler/pkg/common"
)

// CreateEventHandler handler for events
func CreateEventHandler(w http.ResponseWriter, r *http.Request) {
	var eventRequest model.EventRequest
	err := json.NewDecoder(r.Body).Decode(&eventRequest)
	if err != nil {
		common.SendErrorResponse(w, r, common.NewAppError(err.Error(), http.StatusBadRequest))
		return
	}

	createEventErr := internal.NewEventService().CreateEvent(&eventRequest)
	if createEventErr != nil {
		common.SendErrorResponse(w, r, common.NewAppError(createEventErr.Error(), http.StatusInternalServerError))
		return
	}
	common.SendResult(w, r, nil)
	return
}
