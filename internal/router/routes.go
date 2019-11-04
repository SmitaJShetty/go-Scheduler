package router

import (
	"github.com/gorilla/mux"
)

// addRoutes adds routes
func addRoutes(router *mux.Router) {

}

// GetRouter gets router
func GetRouter() *mux.Router {
	router := mux.NewRouter()
	addRoutes(router)
	return router
}
