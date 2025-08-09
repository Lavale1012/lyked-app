package routes

import (
	authHandlers "lyked-backend/internal/handlers/auth"
	testHandlers "lyked-backend/internal/handlers/test"

	"github.com/gin-gonic/gin"
)

func InitUserRoutes(r *gin.Engine) error {
	userRoutes := r.Group("/users")
	{
		userRoutes.GET("/test", testHandlers.TestRoute)
		// Define user-related routes here, e.g.:
		userRoutes.POST("/register", authHandlers.RegisterUser)
		userRoutes.POST("/login", authHandlers.LoginUser)
	}
	return nil
}
