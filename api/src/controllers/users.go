package controllers

import (
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

// CreateUser create a new user
func CreateUser(rw http.ResponseWriter, r *http.Request) {
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.Error(rw, http.StatusUnprocessableEntity, err)
		return
	}

	var user models.User
	if err = json.Unmarshal(requestBody, &user); err != nil {
		responses.Error(rw, http.StatusBadRequest, err)
		return
	}
	if err = user.Prepare(http.MethodPost); err != nil {
		responses.Error(rw, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.Error(rw, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewUserRepository(db)

	user.ID, err = repository.Create(user)
	if err != nil {
		responses.Error(rw, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(rw, http.StatusOK, user)
}

// GetUsers get users filtering by name or login
func GetUsers(rw http.ResponseWriter, r *http.Request) {
	nameOrLogin := strings.ToLower(r.URL.Query().Get("user"))

	db, err := database.Connect()
	if err != nil {
		responses.Error(rw, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewUserRepository(db)

	users, err := repository.GetUsers(nameOrLogin)
	if err != nil {
		responses.Error(rw, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(rw, http.StatusOK, users)
}

// GetUserById use one ID to get an user
func GetUserById(rw http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	userID, err := strconv.ParseUint(parameters["id"], 10, 64)
	if err != nil {
		responses.Error(rw, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.Error(rw, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewUserRepository(db)
	user, err := repository.GetUserById(userID)
	if err != nil {
		responses.Error(rw, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(rw, http.StatusOK, user)
}

// UpdateUser update an user
func UpdateUser(rw http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	userID, err := strconv.ParseUint(parameters["id"], 10, 64)
	if err != nil {
		responses.Error(rw, http.StatusBadRequest, err)
		return
	}

	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.Error(rw, http.StatusBadRequest, err)
		return
	}

	var user models.User
	if err := json.Unmarshal(requestBody, &user); err != nil {
		responses.Error(rw, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.Error(rw, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewUserRepository(db)
	err = repository.UpdateUser(userID, user)
	if err != nil {
		responses.Error(rw, http.StatusInternalServerError, err)
		return
	}

	user, err = repository.GetUserById(userID)
	if err != nil {
		responses.Error(rw, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(rw, http.StatusOK, user)
}

// DeleteUser delete an user
func DeleteUser(rw http.ResponseWriter, r *http.Request) {
	rw.WriteHeader(http.StatusNotImplemented)
}
