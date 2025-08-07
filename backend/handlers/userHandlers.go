package handlers

import (
	"context"
	PDB "lyked-backend/db/postgresql"
	modelPG "lyked-backend/model/posgressModels"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(c *gin.Context) {
	var db = PDB.PostgresDB
	var user modelPG.User

	// First, bind the JSON data
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request data", "details": err.Error()})
		return
	}

	// Validate required fields
	if user.Username == "" || user.Email == "" || user.Password == "" {
		c.JSON(400, gin.H{"error": "Username, email, and password are required"})
		return
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to hash password", "details": err.Error()})
		return
	}
	user.Password = string(hashedPassword)

	// Create user in database
	ctx := context.Background()
	err = db.WithContext(ctx).Create(&user).Error
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to register user", "details": err.Error()})
		return
	}

	// Return success without exposing the password
	c.JSON(201, gin.H{
		"message": "User registered successfully",
		"user": gin.H{
			"id":       user.ID,
			"username": user.Username,
			"email":    user.Email,
		},
	})
}

func LoginUser(c *gin.Context) {}
