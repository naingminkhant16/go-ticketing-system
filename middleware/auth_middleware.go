package middleware

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	jwt2 "ticketing-system/common/jwt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized",
			})
			return
		}

		tokenString := strings.TrimSpace(authHeader)
		if strings.HasPrefix(strings.ToLower(tokenString), "bearer ") {
			tokenString = strings.TrimSpace(tokenString[7:])
		}

		var claims jwt2.Claims
		// TODO : fix Unauthorized response event with token
		token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(jwt2.Secret), nil
		}, jwt.WithLeeway(2*time.Minute))

		log.Println(token, err)
		if err != nil {
			log.Println("JWT Parsing Error:", err)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized",
			})
			return
		}
		if !token.Valid {
			log.Println("JWT Token is parsed but invalid")
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
			return
		}

		// extract claims
		c.Set("user_id", claims.UserID)
		c.Set("email", claims.Email)
		c.Set("role", claims.Role)
		c.Next()
	}
}
