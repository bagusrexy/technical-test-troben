package config

import (
	"log"
	"os"
	"strconv"
)

var (
	ApiKey    string
	RedisAddr string
	RedisPass string
	RedisDB   int
)

func LoadConfig() {
	ApiKey = os.Getenv("API_KEY")
	if ApiKey == "" {
		log.Fatalf("API_KEY is not set")
	}
	RedisAddr = os.Getenv("REDIS_HOST")
	RedisPass = os.Getenv("REDIS_PASS")
	RedisDB = 4
	if db, err := strconv.Atoi(os.Getenv("REDIS_DB")); err == nil {
		RedisDB = db
	}

}
