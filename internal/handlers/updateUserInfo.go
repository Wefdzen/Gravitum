package handlers

import (
	"net/http"

	"github.com/Wefdzen/Gravitum/internal/db"
	"github.com/gin-gonic/gin"
)

type UpdateData struct {
	UserName    string `json:"username"`
	NewNameUser string `json:"newusername"` // The new username to be set
	NewPassword string `json:"newpassword"` // The new password to be set
}

// UpdateUserInfo returns a Gin handler function for updating user information.
// It reads the update data from the JSON request body, validates it, and then
// calls the repository to update the user information in the database.
func UpdateUserInfo() gin.HandlerFunc {
	return func(c *gin.Context) {
		var jsonInput UpdateData
		if err := c.BindJSON(&jsonInput); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		userRepo := db.NewGormUserRepository()
		db.UpdateUserInfo(userRepo, jsonInput.UserName, jsonInput.NewNameUser, jsonInput.NewPassword)
		c.Status(http.StatusOK)
	}
}
