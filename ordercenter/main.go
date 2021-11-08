package main

import (
	_ "ordercenter/common/configinfo"

	_ "ordercenter/common/cache"
	// _ "ordercenter/application/env/db/initialize"
	_ "ordercenter/common/datasource"
	_ "ordercenter/common/log"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"ordercenter/router"
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