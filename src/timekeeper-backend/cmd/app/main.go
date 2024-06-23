package main

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "timekeeper-backend/docs"
	"timekeeper-backend/internal/db"
	"timekeeper-backend/internal/health"
	"timekeeper-backend/internal/remote"
)

// @title Timekeeper Backend API
// @version 1.0
// @description This is a server for timekeeper backend.
// @host localhost:8080
// @BasePath /
func main() {
	// Initialize the database
	database := db.Init()
	defer database.Close()

	// Set up the Gin router
	r := gin.Default()

	// Health check endpoint
	r.GET("/health-check", health.HealthCheck)

	// Set up routes
	r.GET("/remote-names", remote.GetRemoteNamesHandler(database))
	r.GET("/dashboard", remote.DashboardHandler(database))
	r.GET("/get-remote", remote.GetRemoteHandler(database))
	r.POST("/push-remote", remote.PushRemoteHandler(database))

	// Swagger endpoint
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Run the server
	r.Run(":8080")
}
