package config

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

var envFile = ".env"

func Init() map[string]string {
	params, err := godotenv.Read(envFile)
	if err != nil {
		log.Fatal(fmt.Sprintf("Error loading %s file", envFile))
	}

	return params
}
