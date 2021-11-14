package routers

import (
	"github.com/gin-gonic/gin"
	"usercenter/internal/pkg/cache"
	"xorm.io/xorm"

	"usercenter/internal/adapter/persistence"
	"usercenter/internal/adapter/rest"
	"usercenter/internal/application"
	"usercenter/internal/pkg/middleware"
)

type IRouter interface {
	Init() *gin.Engine
}

type Router struct {
	authContr controller.IAuthController
}

func Build(db *xorm.EngineGroup, rdbRepo cache.ICacheRepository) IRouter {
	mysqlRepo := persistence.BuildMysqlRepo(db)
	return Router {
		authContr: controller.Build(service.Build(mysqlRepo, rdbRepo)),
	}
}

// Init 初始化
func (r Router) Init() *gin.Engine {

	router := gin.Default()
	router.Use(middleware.ErrorHandler())

	r.authContr.InitRouters(router)

	return router
}
