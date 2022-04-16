package routers

import (
	"github.com/gin-gonic/gin"
	"xorm.io/xorm"

	"pjhu/medicine/middleware"
	"pjhu/medicine/pkg/cache"
	"pjhu/medicine/pkg/usercenter/adapter/persistence"
	"pjhu/medicine/pkg/usercenter/adapter/rest"
	"pjhu/medicine/pkg/usercenter/application"
)

type IRouter interface {
	Init() *gin.Engine
}

type Router struct {
	authContr rest.IAuthController
}

func Build(db *xorm.EngineGroup, rdbRepo cache.ICacheRepository) IRouter {
	mysqlRepo := persistence.BuildMysqlRepo(db)
	return Router{
		authContr: rest.Build(application.Build(mysqlRepo, rdbRepo)),
	}
}

// Init 初始化
func (r Router) Init() *gin.Engine {

	router := gin.Default()
	router.Use(middleware.ErrorHandler())

	r.authContr.InitRouters(router)

	return router
}
