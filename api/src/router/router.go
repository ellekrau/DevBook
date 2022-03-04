package router

import (
	"api/src/router/routes"

	"github.com/gorilla/mux"
)

// GenerateRouter will return a router with configured routes
func GenerateRouter() *mux.Router {
	router := mux.NewRouter()
	routes.Configure(router)

	return router
}
