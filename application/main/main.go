package main

import (
	_ "medicine/common/main/configinfo"
	_ "medicine/application/main/resources/db/initialize"
	_ "medicine/common/main/datasource"
	_ "medicine/common/main/log"

	log "github.com/sirupsen/logrus"
)

func init() {
}

func main() {
	// 初始化路由
	routerengine := SetupRouterEngine()
	
	if err := routerengine.Run(); err != nil {
		log.Info("startup service failed, err:%v", err)
	}
}