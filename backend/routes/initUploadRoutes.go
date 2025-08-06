package routes

import (
	"lyked-backend/handlers"

	"github.com/gin-gonic/gin"
)

func InitUploadRoutes(r *gin.Engine) error {
	uploadRoutes := r.Group("/uploads")
	{

		uploadRoutes.POST("/upload", handlers.UploadHandler)
		uploadRoutes.DELETE("/delete", handlers.DeleteUploadHandler)
		uploadRoutes.GET("/all", handlers.GetAllUploadsHandler)
		uploadRoutes.GET("/debug", handlers.DebugUploadsHandler)

	}
	return nil
}
