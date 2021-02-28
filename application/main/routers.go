package main

import (
	"github.com/gin-gonic/gin"

	order "medicine/order/main/adapter/rest"
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

	orderGroup := engine.Group("/api/v1/user")

	for _, opt := range include(order.Routers) {
		opt(orderGroup)
	}
	return engine
}