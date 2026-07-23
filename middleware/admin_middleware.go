package middleware

import (
	"net/http"
	"strings"
	"ticketing-system/entity"

	"github.com/gin-gonic/gin"
)

func AdminHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, uExist := c.Get("user_id")
		role, rExist := c.Get("role")

		if !uExist || !rExist {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized",
			})
			return
		}

		roleStr, ok := role.(string)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized",
			})
			return
		}

		roleStr = strings.TrimSpace(roleStr)

		allowedRoles := map[string]bool{
			string(entity.SystemAdmin): true,
			string(entity.Admin):       true,
			string(entity.SuperAdmin):  true,
		}

		if !allowedRoles[roleStr] {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"message": "Forbidden",
			})
			return
		}

		c.Next()
	}
}
