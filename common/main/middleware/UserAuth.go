package middleware

import (
	"net/http"
	"strings"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"

	"medicine/common/main/errors"
	"medicine/common/main/cache"
)

// AuthUserKey is the cookie name for user credential in basic auth.
const AuthUserKey = "user"

// UserAuth returns a Basic HTTP Authorization middleware
// the key is the user name and the value is the password.
func UserAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		validateAndRefreshToken(ctx)
		ctx.Next()
	}
}

func validateAndRefreshToken(ctx *gin.Context) {
	fullTokenString := ctx.Request.Header.Get(cache.AuthorizationHeader)
	if !strings.HasPrefix(fullTokenString, cache.MiniProgramTokenPrefix) {
		ctx.Header("Content-Type", "application/json")
		errWithCode := errors.NewErrorWithCode(errors.SystemInternalError, "prefix error")
		ctx.AbortWithError(http.StatusBadRequest, errWithCode)
		return
	}

	tokenString, err := cache.ExtractTokenKey(fullTokenString)
	if err != nil {
		log.Error(err)
		ctx.Header("Content-Type", "application/json")
		errWithCode := errors.NewErrorWithCode(errors.SystemInternalError, err.Error())
		ctx.AbortWithError(http.StatusBadRequest, errWithCode)
		return
	}

	var userMeta cache.UserMeta
	err = cache.GetBy(cache.UserAuthNameSpace, tokenString, &userMeta)
	if err != nil {
		log.Error(err)
		ctx.Header("Content-Type", "application/json")
		errWithCode := errors.NewErrorWithCode(errors.SystemInternalError, err.Error())
		ctx.AbortWithError(http.StatusBadRequest, errWithCode)
		return
	}
	ctx.Set(AuthUserKey, userMeta)
}