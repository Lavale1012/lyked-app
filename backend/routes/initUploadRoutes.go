package routes

import (
	debugHandlers "lyked-backend/internal/handlers/test"

	"github.com/gin-gonic/gin"
)

func InitUploadRoutes(r *gin.Engine) error {
	uploadRoutes := r.Group("/uploads")
	{

		uploadRoutes.GET("/debug", debugHandlers.DebugUploadsHandler)

	}
	return nil
}
