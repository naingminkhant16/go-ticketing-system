package handler

import (
	"net/http"
	"ticketing-system/common/response"
	"ticketing-system/entity"
	"ticketing-system/entity/dto"
	"ticketing-system/service"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authService *service.AuthService
}

func NewAuthHandler(authService *service.AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

func (ah *AuthHandler) Register(ctx *gin.Context) {
	var input dto.UserRegisterDto

	if err := ctx.ShouldBind(&input); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	user, err := ah.authService.Register(input, entity.Customer, ctx)
	if err != nil {
		ctx.Error(err)
		return
	}

	response.Created(ctx, user)
}
