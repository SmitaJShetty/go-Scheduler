package router

import (
	"github.com/gorilla/mux"
	"github.com/smitajshetty/go-scheduler/internal/handlers"
)

// addRoutes adds routes
func addRoutes(router *mux.Router) {
	router.HandleFunc("/event", handlers.CreateEventHandler).Methods("POST")
}

// GetRouter gets router
func GetRouter() *mux.Router {
	router := mux.NewRouter()
	addRoutes(router)
	return router
}
