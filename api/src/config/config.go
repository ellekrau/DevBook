package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	// ConnectionString from database
	ConnectionString = ""

	// ApiPort application port
	ApiPort = 0
)

func Load() {
	// Load the .env variables to os.env
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	loadApiConfiguration()
	loadDbConfiguration()
}

func loadApiConfiguration() {
	var err error

	ApiPort, err = strconv.Atoi(getEnv("API_PORT"))
	if err != nil {
		log.Fatal(fmt.Sprintf("Error in parse enviroment variable API_PORT to int: %s", err.Error()))
	}
}

func loadDbConfiguration() {
	ConnectionString = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True?loc=Local",
		getEnv("DB_USERNAME"),
		getEnv("DB_PASSWORD"),
		getEnv("DB_NAME"))
}

func getEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatal(fmt.Sprintf("Error in load enviroment variable: %s", key))
	}
	return value
}
