package middlewares

import (
	"api/src/auth"
	"api/src/controllers/responses"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// ValidateUserIDTokenAuthorization validates if the userID from the request is the same of the authorization token
func ValidateUserIDTokenAuthorization(nextFunction http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		fmt.Println("Method", r.Method)
		if r.Method == http.MethodGet {
			nextFunction(rw, r)
		}

		parameters := mux.Vars(r)
		requestUserID, err := strconv.ParseUint(parameters["userID"], 10, 64)
		if err != nil {
			responses.Error(rw, http.StatusBadRequest, err)
			panic(err)
			// return
		}

		tokenUserID, err := auth.GetUserIDFromToken(r)
		if err != nil {
			responses.Error(rw, http.StatusBadRequest, err)
			panic(err)
			// return
		}

		fmt.Printf("TokenUserID: %d | RequestUserID: %d/n", tokenUserID, requestUserID)
		if tokenUserID != requestUserID {
			responses.Error(rw, http.StatusUnauthorized, nil)
			return
		}

		nextFunction(rw, r)
	}
}

// Auth verifies if the user is authenticated
func Auth(nextFunction http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		if err := auth.ValidateToken(r); err != nil {
			responses.Error(rw, http.StatusUnauthorized, err)
			return
		}

		fmt.Println(r.Method)
		if r.Method == http.MethodGet {
			nextFunction(rw, r)
			return
		}

		parameters := mux.Vars(r)
		requestUserID, err := strconv.ParseUint(parameters["userID"], 10, 64)
		if err != nil {
			responses.Error(rw, http.StatusBadRequest, err)
			return
		}

		tokenUserID, err := auth.GetUserIDFromToken(r)
		if err != nil {
			responses.Error(rw, http.StatusBadRequest, err)
			return
		}

		if tokenUserID != requestUserID {
			responses.Error(rw, http.StatusUnauthorized, nil)
			return
		}

		nextFunction(rw, r)
	}
}

// Logger log
func Logger(nextFunction http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		fmt.Printf("%s %s %s \n", r.Method, r.RequestURI, r.Host)
		nextFunction(rw, r)
	}
}
