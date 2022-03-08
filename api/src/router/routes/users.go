package routes

import (
	"api/src/controllers"
	"fmt"
	"net/http"
)

var uriBase string = "/users"
var uriWithId string = fmt.Sprintf("%s/{userID}", uriBase)

var userRoutes = []Route{
	{URI: uriBase, Method: http.MethodPost, Function: controllers.CreateUser, RequiresAuthentication: false},
	{URI: uriBase, Method: http.MethodGet, Function: controllers.GetUsers, RequiresAuthentication: true},
	{URI: uriWithId, Method: http.MethodGet, Function: controllers.GetUserById, RequiresAuthentication: true},
	{URI: uriWithId, Method: http.MethodPut, Function: controllers.UpdateUser, RequiresAuthentication: true},
	{URI: uriWithId, Method: http.MethodDelete, Function: controllers.DeleteUser, RequiresAuthentication: true},
}
