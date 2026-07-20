package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type LoginResponse struct {
	AccessToken  string
	RefreshToken string
}

func Success(c *gin.Context, message string, data any) {
	if message == "" {
		message = "Success"
	}
	c.JSON(http.StatusOK, Response{
		Message: message,
		Data:    data,
	})
}

func Created(c *gin.Context, data any) {
	c.JSON(http.StatusCreated, Response{
		Message: "Created successfully",
		Data:    data,
	})
}

func NoContent(c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		Message: "Success",
	})
}

func NotFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, Response{
		Message: "Not Found",
	})
}
