package auth

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var JWTSecret string = os.Getenv("JWT_SECRET")

type Claims struct {
	UserID string `json:"user_id"`
	Email  string `json:"email"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

type TokenType string

const (
	AccessToken     TokenType = "access_token"
	RefreshToken    TokenType = "refresh_token"
	MailVerifyToken TokenType = "mail_verify_token"
)

func GenerateToken(userID string, email string, role string, tokenType TokenType) (string, error) {
	var expiresAt time.Time

	now := time.Now().UTC()
	switch tokenType {
	case RefreshToken:
		expiresAt = now.Add(730 * time.Hour) // 1 month
	case AccessToken:
		expiresAt = now.Add(1 * time.Hour) // 1 hr
	case MailVerifyToken:
		expiresAt = now.Add(24 * time.Hour) // 24 hr
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

	tokenString, err := token.SignedString([]byte(JWTSecret))

	if err != nil {
		log.Println("Failed to generate token:", err)
		return "", err
	}

	return tokenString, nil
}

func ParseToken(tokenString string) (*Claims, error) {
	var claims Claims
	token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(JWTSecret), nil
	}, jwt.WithLeeway(2*time.Minute))

	if err != nil || !token.Valid {
		log.Println("JWT Parsing Error:", err)
		return nil, err
	}

	return &claims, nil
}
