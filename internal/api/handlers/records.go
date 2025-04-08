package handlers

import (
	"strconv"

	"github.com/Jeecis/goapi/internal/models"
	"github.com/Jeecis/goapi/internal/repository"
	"github.com/gin-gonic/gin"
)

func CreateRecordHandler(recordRepo *repository.RecordRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		var record models.Record
		if err := c.ShouldBindJSON(&record); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		if err := recordRepo.Create(&record); err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(201, record)
	}
}

func GetRecordListHandler(recordRepo *repository.RecordRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		records, err := recordRepo.GetAll()
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, records)
	}
}

func GetRecordHandler(recordRepo *repository.RecordRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		record, err := recordRepo.GetByID(id)
		if err != nil {
			c.JSON(404, gin.H{"error": "Record not found"})
			return
		}

		c.JSON(200, record)
	}
}

func UpdateRecordHandler(recordRepo *repository.RecordRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		var record models.Record
		if err := c.ShouldBindJSON(&record); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		if err := recordRepo.Update(&record); err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, record)
	}
}

func DeleteRecordHandler(recordRepo *repository.RecordRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		idUint, err := strconv.ParseUint(id, 10, 32)
		if err != nil {
			c.JSON(400, gin.H{"error": "Invalid ID format"})
			return
		}

		if err := recordRepo.Delete(uint(idUint)); err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, gin.H{"message": "Record deleted successfully"})
	}
}
