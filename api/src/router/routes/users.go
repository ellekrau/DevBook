package routes

import (
	"fmt"
	"net/http"
)

var uriBase string = "/users"
var uriWithId string = fmt.Sprintf("%s/{id}", uriBase)

var userRoutes = []Route{
	{URI: uriBase, Method: http.MethodPost, Function: create, RequiresAuthentication: false},
	{URI: uriBase, Method: http.MethodGet, Function: getAll, RequiresAuthentication: false},
	{URI: uriWithId, Method: http.MethodGet, Function: getById, RequiresAuthentication: false},
	{URI: uriWithId, Method: http.MethodPost, Function: update, RequiresAuthentication: false},
	{URI: uriWithId, Method: http.MethodDelete, Function: delete, RequiresAuthentication: false},
}

func create(rw http.ResponseWriter, r *http.Request)  {}
func getAll(rw http.ResponseWriter, r *http.Request)  {}
func getById(rw http.ResponseWriter, r *http.Request) {}
func update(rw http.ResponseWriter, r *http.Request)  {}
func delete(rw http.ResponseWriter, r *http.Request)  {}
