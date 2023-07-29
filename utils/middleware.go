package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(context *gin.Context) {
		tokenString := context.GetHeader("Authorization")
		if tokenString == "" {
			ErrorResponse(context, http.StatusBadRequest, "request does not contain an access token")
			context.Abort()
			return
		}
		err := ValidateToken(tokenString)
		if err != nil {
			ErrorResponse(context, http.StatusUnauthorized, err.Error())
			context.Abort()
			return
		}
		context.Next()
	}
}
