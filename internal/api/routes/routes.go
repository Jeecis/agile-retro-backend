package routes

import (
	"database/sql"

	"github.com/Jeecis/goapi/internal/api/handlers"
	"github.com/Jeecis/goapi/internal/api/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go"
)

func SetupRouter(db *sql.DB, minio *minio.Client) *gin.Engine {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:8080", "http://localhost:3000"}, // Add allowed origins
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// Health check route
	r.GET("/api/health", handlers.HealthCheck)

	// API v1 routes
	v1 := r.Group("/api/v1")
	{
		// Public routes
		auth := v1.Group("/auth")
		{
			auth.POST("/login", handlers.Login)
			auth.POST("/register", handlers.Register)
		}

		// Protected routes
		protected := v1.Group("/")
		protected.Use(middleware.AuthMiddleware())
		{
			users := protected.Group("/users")
			{
				users.GET("/", handlers.GetUsers)
				users.GET("/:id", handlers.GetUser)
				users.PUT("/:id", handlers.UpdateUser)
				users.DELETE("/:id", handlers.DeleteUser)
			}

			resources := protected.Group("/resources")
			{
				resources.GET("/", handlers.GetResourceList)
				resources.GET("/:id", handlers.GetResource)
				resources.POST("/", handlers.AddResource)
				resources.DELETE("/:id", handlers.DeleteResource)
			}

			// Add other resource routes...
		}
	}

	return r
}
