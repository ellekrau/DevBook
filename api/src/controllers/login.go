package controllers

import (
	"api/src/auth"
	"api/src/controllers/responses"
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"api/src/security"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

// Login user authentication in the API
func Login(rw http.ResponseWriter, r *http.Request) {
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

	db, err := database.Connect()
	if err != nil {
		responses.Error(rw, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewUserRepository(db)

	dbUser, err := repository.GetUserByEmail(user.Email)
	if err != nil {
		responses.Error(rw, http.StatusInternalServerError, err)
		return
	}

	if err := security.VerifyPassword(dbUser.Password, user.Password); err != nil {
		responses.Error(rw, http.StatusUnauthorized, errors.New("wrong login or password"))
		return
	}

	token, err := auth.CreateToken(dbUser.ID)
	if err != nil {
		responses.Error(rw, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(rw, http.StatusOK, token)
}
