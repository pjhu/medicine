package controller

import (
	"github.com/sirupsen/logrus"
	"net/http"

	"github.com/gin-gonic/gin"
	"usercenter/internal/application"
	"usercenter/internal/application/command"
	"usercenter/internal/pkg/cache"
)

type IAuthController interface {
	InitRouters(router *gin.Engine)
}

type AuthController struct {
	appSvc service.IApplicationService
}

func Build(app service.IApplicationService) IAuthController {
	return AuthController {
		appSvc: app,
	}
}

// InitRouters for order
func (ac AuthController) InitRouters(router *gin.Engine) {

	router.POST("/api/v1/customer/signin", ac.signin)
	router.POST("/api/v1/customer/signout", ac.signout)
	router.POST("/internal-api/v1/varify-token", ac.validateToken)
}

// signin for user authentication
func (ac AuthController) signin (ctx *gin.Context) {
	var signinCommand command.SigninCommand
	if err := ctx.ShouldBindJSON(&signinCommand); err != nil {

		logrus.Error(err)
		ctx.Header("Content-Type", "application/json")
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H {
			"data": nil,
			"message": err.Error(),
		})
		return
	}
	response, err:= ac.appSvc.Signin(signinCommand)
	if err != nil {
		logrus.Error(err)
		ctx.Header("Content-Type", "application/json")
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H {
			"data": nil,
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H {
		"data": response,
		"message": nil,
	})
}

// signout for user authentication
func (ac AuthController) signout (ctx *gin.Context) {
	err := ac.appSvc.Signout(ctx.GetHeader(cache.AuthorizationHeader))
	if err != nil {
		logrus.Error(err)
		ctx.Header("Content-Type", "application/json")
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H {
			"data": nil,
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H {
		"data": "ok",
		"message": nil,
	})
}

// validateToken for user authentication
func (ac AuthController) validateToken (ctx *gin.Context) {
	var validateTokenCommand command.ValidateTokenCommand
	if err := ctx.ShouldBindJSON(&validateTokenCommand); err != nil {
		logrus.Error(err)
		ctx.Header("Content-Type", "application/json")
		ctx.AbortWithStatusJSON(http.StatusBadRequest,  gin.H {
			"data": nil,
			"message": err.Error(),
		})
		return
	}

	userMeta, err := ac.appSvc.ValidateToken(validateTokenCommand.Token)
	if err != nil {
		logrus.Error(err)
		ctx.Header("Content-Type", "application/json")
		ctx.AbortWithStatusJSON(err.GetHTTPStatus(),  gin.H {
			"data": nil,
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H {
		"data": userMeta,
		"message": nil,
	})
}

