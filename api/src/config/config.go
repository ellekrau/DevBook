package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	// DbConnectionString from database
	DbConnectionString = ""

	// ApiPort application port
	ApiPort = 0

	// SecretKey is used to sign the jwt token
	SecretKey []byte
)

// LoadEnviromentVariables loads enviroment variables from .env file
func LoadEnviromentVariables() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	loadApiConfiguration()
	loadDbConfiguration()
	loadSecretKey()
}

// loadApiConfiguration loads the api configuration
func loadApiConfiguration() {
	var err error

	ApiPort, err = strconv.Atoi(getEnv("API_PORT"))
	if err != nil {
		log.Fatal(fmt.Sprintf("Error in parse enviroment variable API_PORT to int: %s", err.Error()))
	}
}

// loadDbConfiguration loads the DB configuration
func loadDbConfiguration() {
	DbConnectionString = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=true&loc=Local",
		getEnv("DB_USERNAME"),
		getEnv("DB_PASSWORD"),
		getEnv("DB_NAME"))
}

// loadSecretKey loads the secret key
func loadSecretKey() {
	SecretKey = []byte(getEnv("SECRET_KEY"))
}

// getEnv returns the value of a variable in the env file
func getEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatal(fmt.Sprintf("Error in load enviroment variable: %s", key))
	}
	return value
}
