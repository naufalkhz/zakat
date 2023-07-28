package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(context *gin.Context) {
		tokenString := context.GetHeader("Authorization")
		if tokenString == "" {
			SendResponse(context, http.StatusBadRequest, "request does not contain an access token")
			context.Abort()
			return
		}
		err := ValidateToken(tokenString)
		if err != nil {
			SendResponse(context, http.StatusBadRequest, err.Error())
			context.Abort()
			return
		}
		context.Next()
	}
}
