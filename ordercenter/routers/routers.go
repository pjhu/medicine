package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	"xorm.io/xorm"

	"ordercenter/internal/adapter/persistence"
	"ordercenter/internal/adapter/rest"
	"ordercenter/internal/application"
	"ordercenter/internal/pkg/middleware"
)

type IRouter interface {
	Init() *gin.Engine
}

type Router struct {
	orderContr controller.IOrderController
}

func Build(db *xorm.EngineGroup, restClient *resty.Client) IRouter {
	repo := persistence.BuildMysqlRepo(db)
	return Router {
		orderContr: controller.Build(service.Build(repo, restClient)),
	}
}

// Init 初始化
func (r Router) Init() *gin.Engine {

	router := gin.Default()
	router.Use(middleware.ErrorHandler())

	r.orderContr.InitRouters(router)

	return router
}
