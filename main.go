package main

import (
	"log"

	"github.com/bagusrexy/technical-test-troben/router"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalln(err)
	}
	r := gin.Default()
	router.Router(r)
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Unable to start server: ", err)
	}
}
