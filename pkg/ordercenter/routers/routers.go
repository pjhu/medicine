package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	"xorm.io/xorm"

	"pjhu/medicine/middleware"
	"pjhu/medicine/pkg/cache"
	"pjhu/medicine/pkg/ordercenter/adapter/persistence"
	"pjhu/medicine/pkg/ordercenter/adapter/rest"
	"pjhu/medicine/pkg/ordercenter/application"
)

type IRouter interface {
	Init() *gin.Engine
}

type Router struct {
	orderContr rest.IOrderController
}

func Build(db *xorm.EngineGroup, rdbRepo cache.ICacheRepository, restClient *resty.Client) IRouter {
	repo := persistence.BuildMysqlRepo(db)
	return Router{
		orderContr: rest.Build(application.Build(repo, restClient)),
	}
}

// Init 初始化
func (r Router) Init() *gin.Engine {

	router := gin.Default()
	router.Use(middleware.ErrorHandler())

	r.orderContr.InitRouters(router)

	return router
}
