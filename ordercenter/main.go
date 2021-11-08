package main

import (
	_ "ordercenter/common/configinfo"
	// _ "ordercenter/application/resources/db/initialize"
	_ "ordercenter/common/datasource"
	_ "ordercenter/common/log"
	_ "ordercenter/common/cache"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	router "ordercenter/router"
)

func init() {
}

func main() {
	// 初始化路由
	routerengine := router.SetupRouterEngine()
	
	if err := routerengine.Run(viper.GetString("gin.port")); err != nil {
		log.WithError(err).Error("startup service failed")
	}
}