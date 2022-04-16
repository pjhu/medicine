package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"

	"pjhu/medicine/middleware"
)

// SetupRouterEngine 初始化
func SetupRouterEngine() *gin.Engine {

	engine := gin.Default()

	engine.Use(middleware.ErrorHandler())
	customerGroup := engine.Group("/api/v1/")
	customerGroup.GET("/*orderPath", middleware.UserAuth(),
		middleware.ReverseProxy(viper.GetString("microservice.ordercenter"), "orderPath"))
	customerGroup.POST("/*orderPath", middleware.UserAuth(),
		middleware.ReverseProxy(viper.GetString("microservice.ordercenter"), "orderPath"))
	return engine
}