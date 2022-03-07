package middlewares

import (
	"fmt"
	"net/http"
)

// Auth verifies if the user is authenticated
func Auth(next http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		fmt.Println("Auth")
		next(rw, r)
	}
}

// Logger log
func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		fmt.Printf("%s %s %s \n", r.Method, r.RequestURI, r.Host)
		next(rw, r)
	}
}
