package middleware

import (
	"github.com/gin-gonic/gin"
)

// AuthUserKey is the cookie name for user credential in basic auth.
const AuthUserKey = "user"

// UserAuth returns a Basic HTTP Authorization middleware
// the key is the username and the value is the password.
func UserAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		validateAndRefreshToken(ctx)
		ctx.Next()
	}
}

func validateAndRefreshToken(ctx *gin.Context) {
}