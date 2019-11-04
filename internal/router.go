package internal

import (
	"fmt"
	"net/http"
)

// Start starts listener
func Start(listenAddress string) {
	router := GetRouter()
	addRoutes(router)
	go func() {
		err := http.ListenAndServe(listenAddress, router)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("Listening on port:", listenAddress)
	}()
}



// addRoutes adds routes
func addRoutes(router *mux.Router) {

}

// GetRouter gets router
	func GetRouter()  *mux.Router {
	router := mux.NewRouter()
	addRoutes(router)
	return router
}