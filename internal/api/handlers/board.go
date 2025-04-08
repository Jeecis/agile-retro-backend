package handlers

import (
	"strconv"

	"github.com/Jeecis/goapi/internal/models"
	"github.com/Jeecis/goapi/internal/repository"
	"github.com/gin-gonic/gin"
)

func CreateBoardHandler(boardRepo *repository.BoardRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		var board models.Board
		if err := c.ShouldBindJSON(&board); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		if err := boardRepo.Create(&board); err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(201, board)
	}
}

func GetBoardListHandler(boardRepo *repository.BoardRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		boards, err := boardRepo.GetAll()
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, boards)
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

func DeleteBoardHandler(boardRepo *repository.BoardRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		idUint, err := strconv.ParseUint(id, 10, 32)
		if err != nil {
			c.JSON(400, gin.H{"error": "Invalid ID format"})
			return
		}

		if err := boardRepo.Delete(uint(idUint)); err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, gin.H{"message": "Board deleted successfully"})
	}
}
