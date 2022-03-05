package main

import (
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	router := router.GenerateRouter

	fmt.Println("API is running!")
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", "5000"), router()))
}
