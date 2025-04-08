package handlers

import (
	"strconv"

	"github.com/Jeecis/goapi/internal/models"
	"github.com/Jeecis/goapi/internal/repository"
	"github.com/gin-gonic/gin"
)

func CreateColumnHandler(columnRepo *repository.ColumnRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		var column models.Column
		if err := c.ShouldBindJSON(&column); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		if err := columnRepo.Create(&column); err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(201, column)
	}
}

func GetColumnListHandler(columnRepo *repository.ColumnRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		columns, err := columnRepo.GetAll()
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, columns)
	}
}

func GetColumnHandler(columnRepo *repository.ColumnRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		column, err := columnRepo.GetByID(id)
		if err != nil {
			c.JSON(404, gin.H{"error": "Column not found"})
			return
		}

		c.JSON(200, column)
	}
}

func UpdateColumnHandler(columnRepo *repository.ColumnRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		var column models.Column
		if err := c.ShouldBindJSON(&column); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		if err := columnRepo.Update(&column); err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, column)
	}
}

func DeleteColumnHandler(columnRepo *repository.ColumnRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		idUint, err := strconv.ParseUint(id, 10, 32)
		if err != nil {
			c.JSON(400, gin.H{"error": "Invalid ID format"})
			return
		}

		if err := columnRepo.Delete(uint(idUint)); err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, gin.H{"message": "Column deleted successfully"})
	}
}
