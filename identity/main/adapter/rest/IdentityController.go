package identitycontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	identitycommand "medicine/identity/main/application/command"
	identityapplicationservice "medicine/identity/main/application/services"
)

// Signin for user authentication
func Signin (ctx *gin.Context) {
	var signinCommand identitycommand.SignoutCommand
	if err := ctx.ShouldBindJSON(&signinCommand); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	response, err:= identityapplicationservice.Signin(signinCommand)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, response)
}

// Signout for user authentication
func Signout (ctx *gin.Context) {
	identityapplicationservice.Signout(ctx.GetHeader("Authorization"))
	ctx.JSON(http.StatusOK, gin.H{"data": "ok"})
}
