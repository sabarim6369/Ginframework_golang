package controllers

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

func CreateUser(c *gin.Context) {
	var user User

	data, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format"})
		return
	}

	fmt.Println("Raw JSON:", string(data))

	err = json.Unmarshal(data, &user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}

	if user.Name == "" || user.Email == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Both 'name' and 'email' are required",
		})
		return
	}

	fmt.Println("Parsed Struct:", user)

	c.JSON(http.StatusOK, gin.H{
		"message": "User created successfully",
		"name":    user.Name,
		"email":   user.Email,
	})
}
