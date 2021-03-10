package main

import (
	_ "usercenter/common/main/configinfo"
	_ "usercenter/application/main/resources/db/initialize"
	_ "usercenter/common/main/datasource"
	_ "usercenter/common/main/log"
	_ "usercenter/common/main/cache"

	log "github.com/sirupsen/logrus"
)

func init() {
}

func main() {
	// 初始化路由
	routerengine := SetupRouterEngine()
	
	if err := routerengine.Run(":8081"); err != nil {
		log.WithError(err).Error("startup service failed")
	}
}