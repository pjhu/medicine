package main

import (
	"github.com/gin-gonic/gin"

	identity "medicine/identity/main/adapter/rest"
	order "medicine/order/main/adapter/rest"
	middleware "medicine/common/main/middleware"
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

	engine.POST("/api/v1/customer/signin", identity.Signin)
	engine.POST("/api/v1/customer/signout", identity.Signout)

	customerGroup := engine.Group("/api/v1/customer")
	customerGroup.Use(middleware.UserAuth())

	for _, opt := range include(order.Routers) {
		opt(customerGroup)
	}
	return engine
}