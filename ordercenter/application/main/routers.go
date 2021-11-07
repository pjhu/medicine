package main

import (
	"github.com/gin-gonic/gin"

	order "ordercenter/core/main/adapter/rest"
	middleware "ordercenter/common/main/middleware"
)

type option func(*gin.RouterGroup)

// include 注册app的路由配置
func include(opts ...option) []option {
	var options = []option{}
	options = append(options, opts...)
	return options
}

// SetupRouterEngine 初始化
func SetupRouterEngine() *gin.Engine {

	engine := gin.Default()
	engine.Use(middleware.ErrorHandler())

	customerGroup := engine.Group("/api/v1/customer")
	// customerGroup.Use(middleware.UserAuth())

	for _, opt := range include(order.Routers) {
		opt(customerGroup)
	}
	return engine
}