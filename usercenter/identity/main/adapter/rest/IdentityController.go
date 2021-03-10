package identitycontroller

import (
	"net/http"
	"github.com/gin-gonic/gin"

	"usercenter/common/main/errors"
	"usercenter/common/main/cache"
	identitycommand "usercenter/identity/main/application/command"
	identityapplicationservice "usercenter/identity/main/application/services"
)

// Signin for user authentication
func Signin (ctx *gin.Context) {
	var signinCommand identitycommand.SignoutCommand
	if err := ctx.ShouldBindJSON(&signinCommand); err != nil {
		errWithCode := errors.NewErrorWithCode(errors.SystemInternalError, err.Error())
		ctx.AbortWithError(http.StatusBadRequest, errWithCode)
		return
	}
	response, err:= identityapplicationservice.Signin(signinCommand)
	if err != nil {
		errWithCode := errors.NewErrorWithCode(errors.SystemInternalError, err.Error())
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
		ctx.AbortWithError(err.GetHTTPStatus(), errWithCode)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": "ok"})
}
