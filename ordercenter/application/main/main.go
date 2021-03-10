package main

import (
	_ "ordercenter/common/main/configinfo"
	_ "ordercenter/application/main/resources/db/initialize"
	_ "ordercenter/common/main/datasource"
	_ "ordercenter/common/main/log"
	_ "ordercenter/common/main/cache"

	log "github.com/sirupsen/logrus"
)

func init() {
}

func main() {
	// 初始化路由
	routerengine := SetupRouterEngine()
	
	if err := routerengine.Run(":8082"); err != nil {
		log.WithError(err).Error("startup service failed")
	}
}