package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Route represent all api routes
type Route struct {
	URI                    string
	Method                 string
	Function               func(http.ResponseWriter, *http.Request)
	RequiresAuthentication bool
}

func Configure(router *mux.Router) *mux.Router {
	routes := userRoutes

	for _, route := range routes {
		router.HandleFunc(route.URI, route.Function).Methods(route.Method)
	}
}
