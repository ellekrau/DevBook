package main

import (
	"api/src/router"
	"api/src/utils"
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

func main() {
	loadEnvFile()

	port := utils.GetEnvVariable("API_PORT")
	router := router.GenerateRouter

	fmt.Println("API is running!")
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), router()))
}

func loadEnvFile() {
	if err := godotenv.Load("../.env"); err != nil {
		log.Fatal("Error in env file load")
	}
}
