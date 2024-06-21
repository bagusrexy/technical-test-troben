package config

import (
	"log"
	"os"
)

var ApiKey string

func LoadConfig() {
	ApiKey = os.Getenv("API_KEY")
	if ApiKey == "" {
		log.Fatalf("API_KEY is not set")
	}
}
