package handlers

import "github.com/gin-gonic/gin"

func CreateBoard(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "API is running",
	})
}

func GetBoardList(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "API is running",
	})
}

func GetBoard(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "API is running",
	})
}

func UpdateBoard(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "API is running",
	})
}

func DeleteBoard(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "API is running",
	})
}
