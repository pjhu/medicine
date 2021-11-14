package middleware

import (
	"github.com/gin-gonic/gin"
	"usercenter/internal/pkg/errors"
)

// ErrorHandler unify error handler
func ErrorHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()
		var errorMessages []string
		for _, err := range ctx.Errors {
			errorMessages = append(errorMessages, err.Error())
		}

		if len(errorMessages) > 0 {
			ctx.JSON(-1, ctx.Errors.Last().Err.(* errors.ErrorWithCode).GetErrorResponse())
		}
	}
}
