package handler

import (
	"net/http"
	"ticketing-system/common/response"
	"ticketing-system/service"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService *service.UserService
}

func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (uh *UserHandler) GetAllUser(ctx *gin.Context) {
	users, err := uh.userService.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	response.Success(ctx, "Success", users)
}
