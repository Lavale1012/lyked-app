package routes

import (
	"lyked-backend/handlers"

	"github.com/gin-gonic/gin"
)

func InitUserRoutes(r *gin.Engine) error {
	userRoutes := r.Group("/users")
	{
		userRoutes.GET("/test", handlers.TestRoute)
		// Define user-related routes here, e.g.:
		// userRoutes.POST("/register", handlers.RegisterUser)
		// userRoutes.POST("/login", handlers.LoginUser)
	}
	return nil
}
