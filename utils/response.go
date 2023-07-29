package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Success struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func SuccessResponse(c *gin.Context, status int, data interface{}) {
	c.JSON(status, Success{
		Status:  status,
		Message: http.StatusText(status),
		Data:    data,
	})
}

type Errors struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Errors  interface{} `json:"errors"`
}

func ErrorResponse(c *gin.Context, status int, errors interface{}) {
	c.JSON(status, Errors{
		Status:  status,
		Message: http.StatusText(status),
		Errors:  errors,
	})
}
