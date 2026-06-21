package route

import (
	"ticketing-system/handler"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(router *gin.Engine, userHandler *handler.UserHandler) {
	rg := router.Group("/api/users")
	{
		rg.GET("", userHandler.GetAllUser)
	}
}

func RegisterAuthRoutes(router *gin.Engine, authHandler *handler.AuthHandler) {
	rg := router.Group("/api/auth")
	{
		rg.POST("/register", authHandler.Register)
	}
}
