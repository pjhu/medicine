package main

import (
	_ "medicine/common/main/configinfo"
	_ "medicine/application/main/resources/db/initialize"
	_ "medicine/common/main/datasource"
	_ "medicine/common/main/log"
	_ "medicine/common/main/cache"

	log "github.com/sirupsen/logrus"
)

func init() {
}

func main() {
	// 初始化路由
	routerengine := SetupRouterEngine()
	
	if err := routerengine.Run(); err != nil {
		log.WithError(err).Error("startup service failed")
	}
}