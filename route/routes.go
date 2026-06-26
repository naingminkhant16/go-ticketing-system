package route

import (
	"ticketing-system/handler"
	"ticketing-system/middleware"

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
		rg.POST("/login", authHandler.Login)
		rg.POST("/refresh", authHandler.RefreshToken)
		rg.Use(middleware.AuthHandler()).GET("/profile", authHandler.Profile)
		// TODO :  verify email callback endpoint
	}
}
