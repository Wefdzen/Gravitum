package handlers

import (
	"net/http"

	"github.com/Wefdzen/Gravitum/internal/db"
	"github.com/gin-gonic/gin"
)

// CreateUser returns a Gin handler function for creating a new user.
// It reads the user data from the JSON request body, validates it, and then
// calls the repository to create the user in the database.
func CreateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		var jsonInput db.User
		if err := c.BindJSON(&jsonInput); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		userRepo := db.NewGormUserRepository()
		db.CreateUser(userRepo, &jsonInput)
		c.Status(http.StatusCreated)
	}
}
