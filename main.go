package main

import (
	"go_mongo/models"
	"log"
	controllers "restAPI1/contollers"
	"restAPI1/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Use(gin.Logger())

	models.ConnectDatabase()

	movies := r.Group("/movies")
	{
		movies.POST("/", controllers.CreateMovie)
		movies.PUT("/:id", controllers.UpdateMovie)
		movies.DELETE("/:id", controllers.DeleteMovie)
		movies.DELETE("/", controllers.DeleteAll)
		movies.GET("/", controllers.ListAll)
		movies.GET("/one/:name", controllers.FindMovie)
		movies.GET("/all/:name", controllers.FindMovies)
		movies.POST("/many", controllers.CreateMovies)
	}

	log.Println("Server Started")
	r.Run()
}
