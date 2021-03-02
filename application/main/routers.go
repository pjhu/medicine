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

	// Global middleware
	// Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
	// By default gin.DefaultWriter = os.Stdout
	engine.Use(gin.Logger())

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	engine.Use(gin.Recovery())

	engine.POST("/api/v1/customer/signin", identity.Signin)
	engine.POST("/api/v1/customer/signout", middleware.UserAuth(), identity.Signout)

	customerGroup := engine.Group("/api/v1/customer")
	customerGroup.Use(middleware.UserAuth())

	for _, opt := range include(order.Routers) {
		opt(customerGroup)
	}
	return engine
}