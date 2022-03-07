package routes

import (
	"api/src/middlewares"
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

// Configure add all routes in router
func Configure(router *mux.Router) {
	routes := userRoutes
	routes = append(routes, loginRoute)

	for _, route := range routes {
		function := route.Function
		function = middlewares.Logger(function)

		if route.RequiresAuthentication {
			function = middlewares.Auth(function)
		}

		router.HandleFunc(route.URI, function).Methods(route.Method)
	}
}
