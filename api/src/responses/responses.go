package responses

import (
	"encoding/json"
	"log"
	"net/http"
)

func JSON(rw http.ResponseWriter, statusCode int, data interface{}) {
	rw.WriteHeader(statusCode)

	if err := json.NewEncoder(rw).Encode(data); err != nil {
		log.Fatal(err)
	}
}

func Error(rw http.ResponseWriter, statusCode int, data error) {
	JSON(rw, statusCode, struct{ Error string }{data.Error()})
}
