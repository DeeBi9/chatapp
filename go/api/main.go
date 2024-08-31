package main

import (
	"net/http"

	"github.com/Deepanshuisjod/chatapp/auth"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// CORS middleware
	r.Use(cors.Default())

	r.POST("/", func(c *gin.Context) {
		var data auth.UserInfo

		if err := c.BindJSON(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		isUsernameExists, userData, message := data.Auth()
		if isUsernameExists {
			c.JSON(http.StatusConflict, gin.H{"message": message, "data": userData})
		} else {
			c.JSON(http.StatusOK, gin.H{"message": message, "data": userData})
		}
	})

	r.Run(":8080")
}
