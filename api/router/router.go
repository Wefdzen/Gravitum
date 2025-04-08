package router

import (
	"net/http"

	"github.com/Wefdzen/Gravitum/internal/handlers"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.StaticFS("/swagger-docs", http.Dir("./api/swagger"))
	url := ginSwagger.URL("/swagger-docs/swagger.yml")
	r.GET("/ui-swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	userAccounts := r.Group("/users")
	{
		userAccounts.GET("/", handlers.GetUserInfo())
		userAccounts.PUT("/", handlers.UpdateUserInfo())
		userAccounts.POST("/", handlers.CreateUser())
	}

	return r
}
