package main

import (
	"go-ifsc/handlers"
	"go-ifsc/middlewares"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	// Configure cors to only allow supported HTTP methods
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"GET"}
	config.AllowHeaders = []string{"Content-Type"}
	router.Use(cors.New(config))
	router.NoRoute()

	router.GET("/bank/:ifsc", middlewares.CacheCheck, handlers.GetBank)
	// Listen and serve on 0.0.0.0:8080
	router.Run(":8080")
}
