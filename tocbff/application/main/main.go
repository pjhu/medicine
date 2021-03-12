package main

import (
	_ "bff/common/main/configinfo"
	_ "bff/common/main/log"
	_ "bff/common/main/cache"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func init() {
}

func main() {
	// 初始化路由
	routerengine := SetupRouterEngine()
	if err := routerengine.Run(viper.GetString("gin.port")); err != nil {
		log.WithError(err).Error("startup service failed")
	}
}