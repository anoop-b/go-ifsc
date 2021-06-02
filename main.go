package main

import (
	"go-ifsc/handlers"
	"go-ifsc/middlewares"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	// Configure cors to only allow supported HTTP methods
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"GET"}
	config.AllowHeaders = []string{"Content-Type"}
	server.Use(cors.New(config))
	server.Use(gin.Recovery())
	server.NoRoute()

	apiRoutes:= server.Group("/bank")
	apiRoutes.Use(middlewares.CacheCheck())
	apiRoutes.GET("/:ifsc",handlers.GetBank)
	// Listen and serve on 0.0.0.0:8080
	server.Run(":8080")
}
