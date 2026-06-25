package handler

import (
	"log"
	"net/http"
	"ticketing-system/common/response"
	"ticketing-system/config"
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

func (ah *AuthHandler) Login(ctx *gin.Context) {
	var input dto.UserLoginDto
	if err := ctx.ShouldBind(&input); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	tokens, err := ah.authService.Login(input)
	if err != nil {
		ctx.Error(err)
		return
	}
	setRefreshTokenHTTPOnlyCookie(ctx, tokens.RefreshToken)

	response.Success(ctx, "success", tokens.AccessToken)
}

func (ah *AuthHandler) Profile(ctx *gin.Context) {
	userID := ctx.GetString("user_id")

	user, err := ah.authService.GetProfile(userID)
	if err != nil {
		ctx.Error(err)
		return
	}

	response.Success(ctx, "success", user)
}

func (ah *AuthHandler) RefreshToken(ctx *gin.Context) {
	rt, err := ctx.Cookie("refresh_token")

	if err != nil {
		log.Println("RefreshToken cookie not found")
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Unauthorized"})
		return
	}

	accessToken, err := ah.authService.RefreshToken(rt)
	if err != nil {
		log.Println(err)
		ctx.Error(err)
		return
	}
	
	response.Success(ctx, "success", accessToken)
}

func setRefreshTokenHTTPOnlyCookie(ctx *gin.Context, rt string) {
	appDomain := config.GetEnvOrPanic("APP_DOMAIN")
	appEnv := config.GetEnvOrPanic("APP_ENV")

	// set refresh token in http only cookie
	ctx.SetCookie(
		"refresh_token",
		rt,
		86400,
		"/",
		appDomain,
		appEnv == "production",
		true,
	)
}
