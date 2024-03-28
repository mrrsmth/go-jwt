package controllers

import (
	"go-jwt/database"
	"go-jwt/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Profile(c *gin.Context) {
	// Get the email from the authorization middleware
	email, exists := c.Get("email")
	if !exists {
		c.JSON(401, gin.H{
			"Error": "Unauthorized",
		})
		c.Abort()
		return
	}

	// Query the database for the user
	var user models.User
	result := database.GlobalDB.Where("email = ?", email.(string)).First(&user)

	// If the user is not found, return a 404 status code
	if result.Error == gorm.ErrRecordNotFound {
		c.JSON(404, gin.H{
			"Error": "User Not Found",
		})
		c.Abort()
		return
	}

	// If an error occurs while retrieving the user profile, return a 500 status code
	if result.Error != nil {
		c.JSON(500, gin.H{
			"Error": "Could Not Get User Profile",
		})
		c.Abort()
		return
	}

	// Set the user's password to an empty string
	user.Password = ""

	// Return the user profile with a 200 status code
	c.JSON(200, user)
}
