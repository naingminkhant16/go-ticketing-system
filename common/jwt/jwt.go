package jwt

import (
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var Secret string = os.Getenv("JWT_SECRET")

type Claims struct {
	UserID string `json:"user_id"`
	Email  string `json:"email"`
	jwt.RegisteredClaims
}

type TokenType string

const (
	AccessToken  TokenType = "access_token"
	RefreshToken TokenType = "refresh_token"
)

func GenerateToken(userID string, email string, tokenType TokenType) (string, error) {
	var expiresAt time.Time

	switch tokenType {
	case RefreshToken:
		expiresAt = time.Now().Add(24 * time.Hour)
		break
	case AccessToken:
	default:
		expiresAt = time.Now().Add(5 * time.Minute)
	}

	claims := Claims{
		UserID: userID,
		Email:  email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiresAt),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    os.Getenv("JWT_ISSUER"),
		},
	}

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	)

	tokenString, err := token.SignedString([]byte(Secret))

	if err != nil {
		log.Println("Failed to generate token:", err)
		return "", err
	}

	return tokenString, nil
}
