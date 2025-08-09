package server

import (
	"fmt"
	DB "lyked-backend/internal/database/mongodb"
	PDB "lyked-backend/internal/database/postgresql"

	"lyked-backend/internal/utils"
	"lyked-backend/routes"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitServer() error {
	utils.LoadEnv()
	PORT := utils.GetEnv("PORT", "8084")
	r := gin.Default()

	r.Use(gin.Logger())
	r.Use(cors.New(cors.Config{
		AllowAllOrigins:  true, // Allow all origins for development
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: false, // Must be false when AllowAllOrigins is true
		MaxAge:           12 * time.Hour,
	}))
	r.Use(gin.Recovery())

	if err := routes.InitUserRoutes(r); err != nil {
		return fmt.Errorf("failed to initialize user routes: %w", err)
	}
	if err := routes.InitUploadRoutes(r); err != nil {
		return fmt.Errorf("failed to initialize upload routes: %w", err)
	}
	if err := routes.InitProtectedUploadRoutes(r); err != nil {
		return fmt.Errorf("failed to initialize protected upload routes: %w", err)
	}
	if err := routes.InitProtectedUserRoutes(r); err != nil {
		return fmt.Errorf("failed to initialize protected user routes: %w", err)
	}
	// Connect to MongoDB
	if _, err := DB.ConnectMongo("lyked-app"); err != nil {
		return fmt.Errorf("failed to connect to MongoDB: %w", err)
	}

	if _, err := DB.GetCollection("folders"); err != nil {
		return fmt.Errorf("failed to get 'folders' collection: %w", err)
	}
	if _, err := DB.GetCollection("uploads"); err != nil {
		return fmt.Errorf("failed to get 'uploads' collection: %w", err)
	}
	_, err := PDB.ConnectPostgres()
	if err != nil {
		return fmt.Errorf("failed to connect to PostgreSQL: %w", err)
	}

	// Verify the connection was established
	if PDB.PostgresDB == nil {
		return fmt.Errorf("PostgreSQL connection is nil after initialization")
	}

	fmt.Println("✅ PostgreSQL connection verified")

	fmt.Printf("🚀 Server is running on http://localhost:%s\n", PORT)
	// Start the server

	if err := r.Run("localhost:" + PORT); err != nil {
		return fmt.Errorf("failed to start the server: %w", err)
	}
	return nil
}
