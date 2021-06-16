package main

import (
	"embed"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go-ifsc/handlers"
	"go-ifsc/middlewares"
)

//go:embed Data/*.json
var res embed.FS

func main() {
	server := gin.Default()
	// Configure cors to only allow supported HTTP methods
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowMethods = []string{"GET"}
	corsConfig.AllowHeaders = []string{"Content-Type"}
	server.Use(cors.New(corsConfig))
	server.Use(gin.Recovery())
	server.NoRoute()

	apiRoutes := server.Group("/bank")
	apiRoutes.Use(middlewares.CacheCheck())
	apiRoutes.GET("/:ifsc", handlers.GetBank(&res))
	// Listen and serve on 0.0.0.0:8080
	server.Run(":8080")
}
