package handlers

import "github.com/gin-gonic/gin"

func Login(c *gin.Context) {

}

func Register(c *gin.Context) {

}

func GetUser(c *gin.Context) {

}

func GetUsers(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "GetUsers was called",
	})
}

func DeleteUser(c *gin.Context) {

}

func UpdateUser(c *gin.Context) {

}
