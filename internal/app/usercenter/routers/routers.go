package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/pjhu/medicine/internal/app/usercenter/adapter/rest"
	"github.com/pjhu/medicine/internal/pkg/middleware"
)

// Init 初始化
func Init() *gin.Engine {

	router := gin.Default()
	router.Use(middleware.ErrorHandler())
	rest.InitRouters(router)

	return router
}
