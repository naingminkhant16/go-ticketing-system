package middleware

import (
	"errors"
	"net/http"
	apperror "ticketing-system/common/error"

	"github.com/gin-gonic/gin"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) == 0 {
			return
		}

		err := c.Errors.Last().Err
		var appErr *apperror.Error
		if errors.As(err, &appErr) {
			c.JSON(appErr.Status, gin.H{
				"message": appErr.Message,
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal server error",
		})
	}
}
