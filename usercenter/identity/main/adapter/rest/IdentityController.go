package identitycontroller

import (
	"net/http"
	"github.com/gin-gonic/gin"

	"usercenter/common/main/cache"
	"usercenter/common/main/errors"
	"usercenter/identity/main/application/command"
	"usercenter/identity/main/application/services"
)

// Signin for user authentication
func Signin (ctx *gin.Context) {
	var signinCommand identitycommand.SignoutCommand
	if err := ctx.ShouldBindJSON(&signinCommand); err != nil {
		errWithCode := errors.NewErrorWithCode(errors.SystemInternalError, err.Error())
		ctx.Header("Content-Type", "application/json")
		ctx.AbortWithError(http.StatusBadRequest, errWithCode)
		return
	}
	response, err:= identityapplicationservice.Signin(signinCommand)
	if err != nil {
		errWithCode := errors.NewErrorWithCode(errors.SystemInternalError, err.Error())
		ctx.Header("Content-Type", "application/json")
		ctx.AbortWithError(err.GetHTTPStatus(), errWithCode)
		return
	}
	ctx.JSON(http.StatusCreated, response)
}

// Signout for user authentication
func Signout (ctx *gin.Context) {
	err := identityapplicationservice.Signout(ctx.GetHeader(cache.AuthorizationHeader))
	if err != nil {
		errWithCode := errors.NewErrorWithCode(errors.SystemInternalError, err.Error())
		ctx.Header("Content-Type", "application/json")
		ctx.AbortWithError(err.GetHTTPStatus(), errWithCode)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": "ok"})
}

// Signout for user authentication
func ValidateToken (ctx *gin.Context) {
	var validateTokenCommand identitycommand.ValidateTokenCommand
	if err := ctx.ShouldBindJSON(&validateTokenCommand); err != nil {
		errWithCode := errors.NewErrorWithCode(errors.SystemInternalError, err.Error())
		ctx.Header("Content-Type", "application/json")
		ctx.AbortWithError(http.StatusBadRequest, errWithCode)
		return
	}

	userMeta, err := identityapplicationservice.ValidateToken(validateTokenCommand.Token)
	if err != nil {
		errWithCode := errors.NewErrorWithCode(errors.SystemInternalError, err.Error())
		ctx.Header("Content-Type", "application/json")
		ctx.AbortWithError(err.GetHTTPStatus(), errWithCode)
		return
	}
	ctx.JSON(http.StatusOK, userMeta)
}

