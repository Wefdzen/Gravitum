package router

import (
	"github.com/Wefdzen/Gravitum/internal/handlers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	userAccounts := r.Group("/users")
	{
		userAccounts.GET("/", handlers.GetUserInfo())
		userAccounts.PUT("/", handlers.UpdateUserInfo())
		userAccounts.POST("/", handlers.CreateUser())
	}

	return r
}
