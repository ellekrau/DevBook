package router

import "github.com/gorilla/mux"

// GenerateRouter will return a router with configured routes
func GenerateRouter() *mux.Router {
	return mux.NewRouter()
}
