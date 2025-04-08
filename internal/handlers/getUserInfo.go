package handlers

import (
	"net/http"

	"github.com/Wefdzen/Gravitum/internal/db"
	"github.com/gin-gonic/gin"
)

// GetUserInfo returns a Gin handler function for retrieving user information.
// It reads the user data from the JSON request body, validates it, and then
// calls the repository to fetch the user information from the database.
func GetUserInfo() gin.HandlerFunc {
	return func(c *gin.Context) {
		var jsonInput db.User
		if err := c.BindJSON(&jsonInput); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		userRepo := db.NewGormUserRepository()
		res := db.GetUserInfo(userRepo, jsonInput.UserName)
		c.Status(http.StatusOK)
		c.JSON(http.StatusOK, res)
	}
}
