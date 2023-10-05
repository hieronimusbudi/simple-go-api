package helpers

import (
	"github.com/gin-gonic/gin"
)

type (
	ResponsePayload struct {
		StatusCode int         `json:"status_code"`
		Message    string      `json:"message"`
		Data       interface{} `json:"data,omitempty"`
	}
)

func NewResponse(c *gin.Context, statusCode int, message string, data interface{}) {
	response := ResponsePayload{
		StatusCode: statusCode,
		Message:    message,
		Data:       data,
	}
	c.JSON(statusCode, response)
}
