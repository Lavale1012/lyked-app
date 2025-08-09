package routes

import (
	uploadHandlers "lyked-backend/internal/handlers/upload"
	debugHandlers "lyked-backend/internal/handlers/test"

	"github.com/gin-gonic/gin"
)

func InitUploadRoutes(r *gin.Engine) error {
	uploadRoutes := r.Group("/uploads")
	{

		uploadRoutes.POST("/upload", uploadHandlers.UploadHandler)
		uploadRoutes.DELETE("/delete", uploadHandlers.DeleteUploadHandler)
		uploadRoutes.GET("/all", uploadHandlers.GetAllUploadsHandler)
		uploadRoutes.GET("/debug", debugHandlers.DebugUploadsHandler)

	}
	return nil
}
