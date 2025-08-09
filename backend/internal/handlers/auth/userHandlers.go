package handlers

import (
	"context"
	PDB "lyked-backend/internal/database/postgresql"
	modelPG "lyked-backend/internal/models/postgresql"
	"lyked-backend/internal/utils"

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

func LoginUser(c *gin.Context) {
	var db = PDB.PostgresDB
	var login_req modelPG.LoginData
	var user modelPG.User
	// Bind the JSON data
	if err := c.ShouldBindJSON(&login_req); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request data", "details": err.Error()})
		return
	}

	// Validate required fields
	if login_req.Email == "" || login_req.Password == "" {
		c.JSON(400, gin.H{"error": "Email and password are required"})
		return
	}

	query := db.Where("email = ?", login_req.Email)
	if login_req.Username != "" {
		query = db.Where("username = ?", login_req.Username)
	}

	err := query.First(&user).Error
	if err != nil {
		c.JSON(401, gin.H{"error": "Invalid email/username or password"})
		return
	}

	// Compare the hashed password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(login_req.Password))
	if err != nil {
		c.JSON(401, gin.H{"error": "Invalid email/username or password"})
		return
	}

	// Generate token (assuming you have a function for this)
	token, err := utils.GenerateToken(user.ID.String(), db)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to generate token", "details": err.Error()})
		return
	}

	// Return success with token
	c.JSON(200, gin.H{
		"message": "Login successful",
		"token":   token,
		"user": gin.H{
			"id":       user.ID,
			"username": user.Username,
			"email":    user.Email,
		},
	})
}
