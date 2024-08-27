package main

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type User struct {
	username string `json:"username"`
	password string `json:"password"`
}

func main() {
	r := gin.Default()

	// CORS middleware
	r.Use(cors.Default())

	r.POST("/", func(c *gin.Context) {
		// Handle the form data or JSON here
		var data User
		if err := c.Bind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Success", "data": data})
		fmt.Println(data.username)
	})

	r.Run(":8080")
}
