package handlers

import (
	"context"
	"fmt"
	DB "lyked-backend/internal/database/mongodb"
	PDB "lyked-backend/internal/database/postgresql"
	model "lyked-backend/internal/models/mongodb"
	modelPG "lyked-backend/internal/models/postgresql"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func UploadHandler(c *gin.Context) {
	// Handle file upload logic here
	var upload model.LykedUploads
	var user modelPG.User

	userID, exist := c.Get("user_id")
	if !exist {
		c.JSON(401, gin.H{"error": "Unauthorized: user_id not found in context"})
		return
	}

	if err := c.ShouldBindJSON(&upload); err != nil {
		c.JSON(400, gin.H{"error": "Invalid upload data"})
		return
	}
	upload.UserID = userID.(string)
	if upload.UserID == "" {
		c.JSON(400, gin.H{"error": "User ID is required"})
		return
	}
	if upload.VideoLink == "" {
		c.JSON(400, gin.H{"error": "Video Link is required"})
		return
	}

	// Check if database connection is available
	if PDB.PostgresDB == nil {
		c.JSON(500, gin.H{"error": "Database connection not available"})
		return
	}

	err := PDB.PostgresDB.Where("id = ?", upload.UserID).First(&user).Error
	if err != nil {
		c.JSON(400, gin.H{"error": "User not found", "details": err.Error(), "user_id": upload.UserID})
		return
	}

	upload.ID = primitive.NewObjectID()

	fmt.Printf("Received upload: %#v\n", upload)
	collection, err := DB.GetCollection("uploads")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err != nil {
		c.JSON(500, gin.H{"error": "Database connection error"})
		return
	}

	_, err = collection.InsertOne(ctx, upload)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to save upload"})
		return
	}

	c.JSON(200, gin.H{"message": "Upload successful", "upload_id": upload.ID.Hex()})

}

func DeleteUploadHandler(c *gin.Context) {
	// Handle file deletion logic here
	id := c.Query("id")
	userID, exist := c.Get("user_id")
	if id == "" {
		c.JSON(400, gin.H{"error": "Upload ID is required"})
		return
	}
	if !exist {
		c.JSON(401, gin.H{"error": "Unauthorized: user_id not found in context"})
		return
	}
	if userID == "" {
		c.JSON(400, gin.H{"error": "User ID is required"})
		return
	}
	collection, err := DB.GetCollection("uploads")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err != nil {
		c.JSON(500, gin.H{"error": "Database connection error"})
		return
	}

	_, err = collection.DeleteOne(ctx, primitive.M{"_id": id})
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to delete upload"})
		return
	}

	c.JSON(200, gin.H{"message": "Upload deleted successfully"})
}

func GetAllUploadsHandler(c *gin.Context) {
	userID, exist := c.Get("user_id")
	if !exist {
		c.JSON(400, gin.H{"error": "User ID is required"})
		return
	}
	if userID == "" {
		c.JSON(400, gin.H{"error": "User ID is required"})
		return
	}

	collection, err := DB.GetCollection("uploads")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err != nil {
		c.JSON(500, gin.H{"error": "Database connection error"})
		return
	}

	cursor, err := collection.Find(ctx, primitive.M{"user_id": userID})
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to fetch uploads"})
		return
	}
	defer cursor.Close(ctx)

	var uploads []model.LykedUploads
	if err = cursor.All(ctx, &uploads); err != nil {
		c.JSON(500, gin.H{"error": "Failed to parse uploads"})
		return
	}

	c.JSON(200, gin.H{"uploads": uploads})
}
