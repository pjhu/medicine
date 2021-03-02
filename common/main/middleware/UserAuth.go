package middleware

import (
	cache "medicine/common/main/cache"
	"strings"
	"net/http"
	"github.com/gin-gonic/gin"
)

// AuthUserKey is the cookie name for user credential in basic auth.
const AuthUserKey = "user"

// UserAuth returns a Basic HTTP Authorization middleware
// the key is the user name and the value is the password.
func UserAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		validateToken(ctx)
		ctx.Next()
	}
}

func validateToken(ctx *gin.Context) {
	token := ctx.GetHeader("Authorization")
	if (!strings.HasPrefix(token, "MP")) {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	var userMeta cache.UserMeta
	err := cache.Get(cache.UserAuthNameSpace, token[3:], &userMeta)
	if err != nil {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	ctx.Set(AuthUserKey, userMeta)
}