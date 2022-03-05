package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config.LoadEnviromentVariables()
	router := router.GenerateRouter

	fmt.Println("API is running in port:", config.ApiPort)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.ApiPort), router()))
}
