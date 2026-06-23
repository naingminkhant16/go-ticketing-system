package service

import (
	"errors"
	"log"
	apperror "ticketing-system/common/error"
	"ticketing-system/common/helper"
	"ticketing-system/common/jwt"
	"ticketing-system/entity"
	"ticketing-system/entity/dto"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/wneessen/go-mail"
	"gorm.io/gorm"
)

type AuthService struct {
	userSvc     *UserService
	mailService *SMTPService
}

func NewAuthService(userSvc *UserService, mailService *SMTPService) *AuthService {
	return &AuthService{userSvc: userSvc, mailService: mailService}
}

func (svc *AuthService) Register(dto dto.UserRegisterDto, role entity.UserRole, ctx *gin.Context) (*entity.User, error) {
	// check email exists
	exist, err := svc.userSvc.GetByEmail(dto.Email)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	if exist != nil {
		return nil, apperror.BadRequest("Email already exists")
	}

	// check passwords
	if dto.Password != dto.ConfirmPassword {
		return nil, apperror.BadRequest("Passwords do not match")
	}

	hashedPassword, err := helper.HashPassword(dto.Password)
	if err != nil {
		return nil, err
	}

	user, err := svc.userSvc.Create(entity.User{
		Name:     dto.Name,
		Email:    dto.Email,
		Password: hashedPassword,
		Role:     role,
		Gender:   dto.Gender,
		Dob:      &dto.Dob,
	})
	if err != nil {
		return nil, err
	}

	if err := svc.mailService.Send(
		ctx,
		user.Email,
		"Please Verify the Email",
		"Click Here to verify your email",
		mail.TypeTextPlain,
	); err != nil {
		log.Println(err)
	}

	return user, nil
}

func (svc *AuthService) Login(input dto.UserLoginDto) (string, error) {

	user, err := svc.userSvc.GetByEmail(input.Email)
	if err != nil {
		return "", apperror.BadRequest("Email not found")
	}

	if user.VerifiedAt == nil {
		return "", apperror.BadRequest("Please verify your email")
	}

	// verify passwords
	if !helper.VerifyPassword(input.Password, user.Password) {
		return "", apperror.BadRequest("Incorrect email or password")
	}

	accessToken, err := jwt.GenerateToken(user.ID.String(), user.Email, jwt.AccessToken)
	if err != nil {
		return "", apperror.InternalServer("Internal Server Error")
	}

	//refreshToken, err := jwt.GenerateToken(user.ID.String(), user.Email, jwt.RefreshToken)
	//if err != nil {
	//	return "", apperror.InternalServer("Internal Server Error")
	//}

	return accessToken, nil
}

func (svc *AuthService) GetProfile(userId string) (*entity.User, error) {
	uid, err := uuid.Parse(userId)
	if err != nil {
		return nil, apperror.NotFound("User not found")
	}

	user, err := svc.userSvc.userRepository.GetById(uid)

	if err != nil {
		return nil, apperror.NotFound("User not found")
	}

	return user, nil
}
