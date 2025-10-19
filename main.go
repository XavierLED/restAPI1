package main

import (
	"go_mongo/models"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Defualt()

	r.Use(gin.Logger())

	models.ConnectDatabase()

	log.Println("Server Started")
	r.run()
}
