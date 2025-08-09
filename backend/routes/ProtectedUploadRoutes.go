package routes

import (
	uploadHandlers "lyked-backend/internal/handlers/upload"
	"lyked-backend/middleware"

	"github.com/gin-gonic/gin"
)

func InitProtectedUploadRoutes(r *gin.Engine) error {
	protectedUploadRoutes := r.Group("/upload")
	protectedUploadRoutes.Use(middleware.JWTAuthMiddleware()) // Add your authentication middleware here
	{
		protectedUploadRoutes.POST("/upload", uploadHandlers.UploadHandler)
		protectedUploadRoutes.DELETE("/delete", uploadHandlers.DeleteUploadHandler)
		protectedUploadRoutes.GET("/all", uploadHandlers.GetAllUploadsHandler)
	}
	return nil
}
