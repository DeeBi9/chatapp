package main

import (
	"fmt"
	"net/http"

	"github.com/Deepanshuisjod/chatapp/auth"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// CORS middleware
	r.Use(cors.Default())

	r.POST("/authorization", func(c *gin.Context) {
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

	r.POST("/authentication", func(c *gin.Context) {
		var data auth.UserInfoInput
		fmt.Println("Sigin")
		if err := c.BindJSON(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON request: " + err.Error()})
			return
		}

		isSignedIn, message, err, token := data.Signin()
		fmt.Println(err)
		if isSignedIn {
			c.JSON(http.StatusOK, gin.H{"message": message, "token": token})
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"message": message})
		}
	})

	r.Run(":8080")
}
