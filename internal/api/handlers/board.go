package handlers

import (
	"net/http"

	"github.com/Jeecis/goapi/internal/models"
	"github.com/Jeecis/goapi/internal/repository"
	service "github.com/Jeecis/goapi/internal/services"
	"github.com/gin-gonic/gin"
)

func CreateBoard(boardRepo *repository.BoardRepository, columnRepo *repository.ColumnRepository) gin.HandlerFunc {
	return func(c *gin.Context) {

		var boardRequest struct {
			Name string `json:"name" binding:"required"`
		}

		if err := c.ShouldBindJSON(&boardRequest); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request: " + err.Error()})
			return
		}

		board, err := service.CreateBoard(boardRepo, columnRepo, boardRequest.Name)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error: " + err.Error()})
			return
		}

		c.JSON(201, board)
	}
}

func GetBoardHandler(boardRepo *repository.BoardRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		board, err := boardRepo.GetByID(id)
		if err != nil {
			c.JSON(404, gin.H{"error": "Board not found"})
			return
		}

		c.JSON(200, board)
	}
}

func UpdateBoardHandler(boardRepo *repository.BoardRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		var board models.Board
		if err := c.ShouldBindJSON(&board); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		if err := boardRepo.Update(&board); err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, board)
	}
}
