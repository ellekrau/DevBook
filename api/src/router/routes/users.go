package routes

import (
	"api/src/controllers"
	"fmt"
	"net/http"
)

var uriBase string = "/users"
var uriWithId string = fmt.Sprintf("%s/{id}", uriBase)

var userRoutes = []Route{
	{URI: uriBase, Method: http.MethodPost, Function: controllers.CreateUser, RequiresAuthentication: false},
	{URI: uriBase, Method: http.MethodGet, Function: controllers.GetUsers, RequiresAuthentication: false},
	{URI: uriWithId, Method: http.MethodGet, Function: controllers.GetUserById, RequiresAuthentication: false},
	{URI: uriWithId, Method: http.MethodPut, Function: controllers.UpdateUser, RequiresAuthentication: false},
	{URI: uriWithId, Method: http.MethodDelete, Function: controllers.DeleteUser, RequiresAuthentication: false},
}
