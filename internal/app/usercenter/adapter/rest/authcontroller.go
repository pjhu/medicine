package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"github.com/pjhu/medicine/internal/app/usercenter/application"
	"github.com/pjhu/medicine/internal/pkg/cache"
	"github.com/pjhu/medicine/internal/pkg/datasource"
)

// InitRouters for order
func InitRouters(router *gin.Engine) {

	router.POST("/api/v1/customer/signin", signin)
	router.POST("/api/v1/customer/signout", signout)
	router.POST("/internal-api/v1/varify-token", validateToken)
}

// signin for user authentication
func signin(ctx *gin.Context) {
	var signinCommand application.SigninCommand
	if err := ctx.ShouldBindJSON(&signinCommand); err != nil {

		logrus.Error(err)
		ctx.Header("Content-Type", "application/json")
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"data":    nil,
			"message": err.Error(),
		})
		return
	}

	appSvc := application.Builder(datasource.GetDB(), cache.Repository())
	response, err := appSvc.Signin(signinCommand)
	if err != nil {
		logrus.Error(err)
		ctx.Header("Content-Type", "application/json")
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"data":    nil,
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"data":    response,
		"message": nil,
	})
}

// signout for user authentication
func signout(ctx *gin.Context) {
	appSvc := application.Builder(datasource.GetDB(), cache.Repository())
	err := appSvc.Signout(ctx.GetHeader(cache.AuthorizationHeader))
	if err != nil {
		logrus.Error(err)
		ctx.Header("Content-Type", "application/json")
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"data":    nil,
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data":    "ok",
		"message": nil,
	})
}

// validateToken for user authentication
func validateToken(ctx *gin.Context) {
	var validateTokenCommand application.ValidateTokenCommand
	if err := ctx.ShouldBindJSON(&validateTokenCommand); err != nil {
		logrus.Error(err)
		ctx.Header("Content-Type", "application/json")
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"data":    nil,
			"message": err.Error(),
		})
		return
	}
	appSvc := application.Builder(datasource.GetDB(), cache.Repository())
	userMeta, err := appSvc.ValidateToken(validateTokenCommand.Token)
	if err != nil {
		logrus.Error(err)
		ctx.Header("Content-Type", "application/json")
		ctx.AbortWithStatusJSON(err.GetHTTPStatus(), gin.H{
			"data":    nil,
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data":    userMeta,
		"message": nil,
	})
}
