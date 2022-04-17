package routers

import (
	"github.com/gin-gonic/gin"
	"xorm.io/xorm"

	"pjhu/medicine/internal/app/usercenter/adapter/persistence"
	"pjhu/medicine/internal/app/usercenter/adapter/rest"
	"pjhu/medicine/internal/app/usercenter/application"
	"pjhu/medicine/internal/pkg/cache"
	"pjhu/medicine/internal/pkg/middleware"
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
