package handlers

import (
	"log"

	service "github.com/Jeecis/goapi/internal/services"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {

}

func Register(c *gin.Context) {
	providedStr := "this is encryption test"
	log.Print(providedStr)
	encryptedB64, err := service.Encrypt(providedStr, "abcdefghigklmnop")
	if err != nil {
		log.Fatal(err)
	}
	log.Print(encryptedB64)

	decryptedStr, err := service.Decrypt(encryptedB64, "abcdefghigklmnop")
	if err != nil {
		log.Fatal(err)
	}

	log.Print(decryptedStr)

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
