package main

import (
	"fmt"
	"gin-postgres/controllers"
	"gin-postgres/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVars()
	initializers.ConnectToDatabase()
}

func main() {
	fmt.Println("Hello World")

	router := gin.Default()
	router.GET("/posts", controllers.GetPosts)
	router.GET("/posts/:id", controllers.GetPostsByID)
	router.POST("/posts", controllers.PostCreate)
	router.PUT("/posts/:id", controllers.PostUpdateByID)
	router.DELETE("/posts/:id", controllers.PostDeleteByID)
	router.Run()

	// listen and serve on 0.0.0.0:8080
	/*
		If you don’t explicitly specify a port in router.Run() (e.g., router.Run(":8080")), Gin checks if the PORT environment variable is set.
		If PORT is set, Gin uses its value as the port to bind the server.
	*/

}
