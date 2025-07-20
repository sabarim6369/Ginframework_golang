package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"github.com/gin-gonic/gin"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.POST("/create", func(c *gin.Context) {
		var user User

		// Step 1: Read raw data from request body
		data, err := io.ReadAll(c.Request.Body)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format"})
			return
		}

		// Debug print: raw JSON string
		fmt.Println("Raw JSON:", string(data))

		// Step 2: Unmarshal into struct
		err = json.Unmarshal(data, &user)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
			return
		}

		// Step 3: Validate required fields
		if user.Name == "" || user.Email == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Both 'name' and 'email' are required",
			})
			return
		}

		// Print to server console
		fmt.Println("Parsed Struct:", user)

		// Step 4: Success response
		c.JSON(http.StatusOK, gin.H{
			"message": "User created successfully",
			"name":    user.Name,
			"email":   user.Email,
		})
	})

	r.Run()
}
