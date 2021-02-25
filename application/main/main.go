package main

import (
	_ "medicine/application/main/resources/db/initialize"
	_ "medicine/common/main/configinfo"
	_ "medicine/common/main/datasource"
	_ "medicine/common/main/log"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"

	order "medicine/order/main/adapter/rest"
)

func init() {
}

func main() {
	// 加载多个APP的路由配置
	Include(order.Routers)
	// 初始化路由
	r := Init()
	
	// Global middleware
	// Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
	// By default gin.DefaultWriter = os.Stdout
	r.Use(gin.Logger())

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	r.Use(gin.Recovery())
	if err := r.Run(); err != nil {
		log.Info("startup service failed, err:%v", err)
	}
}