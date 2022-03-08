package responses

import (
	"encoding/json"
	"log"
	"net/http"
)

func JSON(rw http.ResponseWriter, statusCode int, data interface{}) {
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(statusCode)

	if data != nil {
		if err := json.NewEncoder(rw).Encode(data); err != nil {
			log.Fatal(err)
		}
	}
}

func Error(rw http.ResponseWriter, statusCode int, data error) {
	if data != nil {
		JSON(rw, statusCode, struct{ Error string }{data.Error()})
		return
	}

	JSON(rw, statusCode, nil)
}
