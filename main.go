package main

import (
	"log"

	"github.com/bagusrexy/technical-test-troben/config"
	"github.com/bagusrexy/technical-test-troben/repositories"
	"github.com/bagusrexy/technical-test-troben/router"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalln("Error loading .env file: ", err)
	}
	repositories.InitRedis(config.RedisAddr, config.RedisPass, config.RedisDB)

	r := gin.Default()
	router.Router(r)
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Unable to start server: ", err)
	}
}
