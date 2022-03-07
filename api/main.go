package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func init() {
	config.LoadEnviromentVariables()
}

func main() {
	router := router.GenerateRouter

	fmt.Println("API is running in port:", config.ApiPort)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.ApiPort), router()))
}
