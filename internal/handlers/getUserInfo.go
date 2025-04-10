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
		username := c.Query("username")
		userRepo := db.NewGormUserRepository()
		res := db.GetUserInfo(userRepo, username)
		if res.UserName == "" {
			c.Status(http.StatusNotFound)
			return
		}
		c.JSON(http.StatusOK, res)
	}
}
