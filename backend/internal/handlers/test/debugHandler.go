package handlers

import (
	"context"
	"fmt"
	DB "lyked-backend/internal/database/mongodb"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func DebugUploadsHandler(c *gin.Context) {
	userID := c.Query("user_id")
	if userID == "" {
		c.JSON(400, gin.H{"error": "User ID is required"})
		return
	}

	collection, err := DB.GetCollection("uploads")
	if err != nil {
		c.JSON(500, gin.H{
			"error":   "Database connection error",
			"details": err.Error(),
		})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{"user_id": userID})
	if err != nil {
		c.JSON(500, gin.H{
			"error":   "Failed to fetch uploads",
			"details": err.Error(),
		})
		return
	}
	defer cursor.Close(ctx)

	// Get raw documents to see the actual structure
	var rawDocs []bson.M
	if err = cursor.All(ctx, &rawDocs); err != nil {
		c.JSON(500, gin.H{
			"error":   "Failed to parse raw uploads",
			"details": err.Error(),
		})
		return
	}

	fmt.Printf("Raw documents found: %d\n", len(rawDocs))
	for i, doc := range rawDocs {
		fmt.Printf("Doc %d: %+v\n", i, doc)
		if id, exists := doc["_id"]; exists {
			fmt.Printf("  _id type: %T, value: %+v\n", id, id)
		}
	}

	c.JSON(200, gin.H{
		"message": "Debug info logged to console",
		"count":   len(rawDocs),
		"docs":    rawDocs,
	})
}