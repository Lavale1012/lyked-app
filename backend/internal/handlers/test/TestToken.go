package handlers

import (
	PDB "lyked-backend/internal/database/postgresql"
	"lyked-backend/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func TestTokenRoute(c *gin.Context) {
	userID := c.Query("user_id")
	token, err := utils.GenerateToken(userID, PDB.PostgresDB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token})
}
