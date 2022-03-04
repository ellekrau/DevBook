package utils

import (
	"fmt"
	"log"
	"os"
)

func GetEnvVariable(key string) string {
	value := os.Getenv(key)

	if value == "" {
		log.Fatal(fmt.Sprintf("Error in GetEnvVariable: %s", key))
	}

	fmt.Printf("Load env variable: [%s] %s\n", key, value)
	return value
}
