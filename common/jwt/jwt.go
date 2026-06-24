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
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

type TokenType string

const (
	AccessToken  TokenType = "access_token"
	RefreshToken TokenType = "refresh_token"
)

func GenerateToken(userID string, email string, role string, tokenType TokenType) (string, error) {
	var expiresAt time.Time

	now := time.Now().UTC()
	switch tokenType {
	case RefreshToken:
		expiresAt = now.Add(24 * time.Hour)
	case AccessToken:
		expiresAt = now.Add(2 * time.Hour)
	}

	claims := Claims{
		UserID: userID,
		Email:  email,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiresAt),
			IssuedAt:  jwt.NewNumericDate(now),
			Issuer:    os.Getenv("JWT_ISSUER"),
		},
	}
	log.Println("Expires at ", expiresAt)
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
