package main

import (
	"github.com/gin-gonic/gin"

	middleware "bff/common/main/middleware"
)

// SetupRouterEngine 初始化
func SetupRouterEngine() *gin.Engine {

	engine := gin.Default()

	// engine.Use(middleware.ErrorHandler())
	// customerGroup := engine.Group("/api/v1/")
	// customerGroup.Use(middleware.UserAuth())

	engine.Any("/*orderPath", middleware.ReverseProxy("http://ordercenter:8080/", "orderPath"))
	return engine
}