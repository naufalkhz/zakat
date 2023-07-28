package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func SendResponse(c *gin.Context, status int, data interface{}) {
	c.JSON(status, Response{
		Status:  status,
		Message: http.StatusText(status),
		Data:    data,
	})
}
