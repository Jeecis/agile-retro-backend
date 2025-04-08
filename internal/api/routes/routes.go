package routes

import (
	"github.com/Jeecis/goapi/internal/api/handlers"
	"github.com/Jeecis/goapi/internal/repository"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB, minio *minio.Client, boardRepo *repository.BoardRepository,
	columnRepo *repository.ColumnRepository,
	recordRepo *repository.RecordRepository,
) *gin.Engine {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:8080", "http://localhost:3000"}, // Add allowed origins
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	v0 := r.Group("/api/v0")
	{
		v0.GET("/health", handlers.HealthCheck)

		b := v0.Group("/board")
		b.POST("", handlers.CreateBoardHandler(boardRepo))
		b.GET("", handlers.GetBoardListHandler(boardRepo))
		b.GET("/{id}", handlers.GetBoardHandler(boardRepo))
		b.PUT("", handlers.UpdateBoardHandler(boardRepo))
		b.DELETE("", handlers.DeleteBoardHandler(boardRepo))

		c := v0.Group("/column")
		c.POST("", handlers.CreateColumnHandler(columnRepo))
		c.GET("", handlers.GetColumnListHandler(columnRepo))
		c.GET("/{id}", handlers.GetColumnHandler(columnRepo))
		c.PUT("", handlers.UpdateColumnHandler(columnRepo))
		c.DELETE("", handlers.DeleteColumnHandler(columnRepo))

		r := v0.Group("/record")
		r.POST("", handlers.CreateRecordHandler(recordRepo))
		r.GET("", handlers.GetRecordListHandler(recordRepo))
		r.GET("/{id}", handlers.GetRecordHandler(recordRepo))
		r.PUT("", handlers.UpdateRecordHandler(recordRepo))
		r.DELETE("", handlers.DeleteRecordHandler(recordRepo))
	}

	// API v1 routes .. PS in future when AUTh is implemented
	// v1 := r.Group("/api/v1")
	// {
	// 	// Public routes
	// 	auth := v1.Group("/auth")
	// 	{
	// 		auth.POST("/login", handlers.Login)
	// 		auth.POST("/register", handlers.Register)
	// 	}

	// 	// Protected routes
	// 	protected := v1.Group("/")
	// 	protected.Use(middleware.AuthMiddleware())
	// 	{
	// 		users := protected.Group("/users")
	// 		{
	// 			users.GET("/", handlers.GetUsers)
	// 			users.GET("/:id", handlers.GetUser)
	// 			users.PUT("/:id", handlers.UpdateUser)
	// 			users.DELETE("/:id", handlers.DeleteUser)
	// 		}

	// 		resources := protected.Group("/resources")
	// 		{
	// 			resources.GET("/", handlers.GetResourceList)
	// 			resources.GET("/:id", handlers.GetResource)
	// 			resources.POST("/", handlers.AddResource)
	// 			resources.DELETE("/:id", handlers.DeleteResource)
	// 		}

	// 		// Add other resource routes...
	// 	}
	// }

	return r
}
