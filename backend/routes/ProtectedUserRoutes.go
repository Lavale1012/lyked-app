package routes

import (
	handlers "lyked-backend/internal/handlers/auth"
	"lyked-backend/middleware"

	"github.com/gin-gonic/gin"
)

func InitProtectedUserRoutes(r *gin.Engine) error {
	protectedUserRoutes := r.Group("/users")
	protectedUserRoutes.Use(middleware.JWTAuthMiddleware()) // Add your authentication middleware here
	{
		protectedUserRoutes.POST("/logout", handlers.LogoutUser)
		protectedUserRoutes.POST("/refresh-token", handlers.RefreshToken)
	}
	return nil
}
