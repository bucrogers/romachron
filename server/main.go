package main

import (
	"server/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	port := "8080"

	router := gin.New()
	router.Use(gin.Logger())

	router.Use(cors.Default())

	// these are the endpoints
	router.GET("/romadatetime", routes.GetRomaDateTime)

	//this runs the server and allows it to listen to requests.
	router.Run(":" + port)
}
