package routes

import (
	"github.com/Jeecis/goapi/internal/api/handlers"
	"github.com/Jeecis/goapi/internal/repository"
	"github.com/Jeecis/goapi/internal/ws"
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
		AllowOrigins:     []string{"*"}, // Add allowed origins
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	r.GET("/health", handlers.HealthCheck)
	v1 := r.Group("/api/v1")
	{

		b := v1.Group("/board")
		b.POST("", handlers.CreateBoard(boardRepo, columnRepo)) //simply create new board with now ws connection
		b.GET("/:id/ws", ws.JoinBoard(boardRepo, columnRepo, recordRepo))
		b.DELETE("", handlers.DeleteBoardHandler(boardRepo)) //delete board with all records and disconnect all users

	}

	return r
}
