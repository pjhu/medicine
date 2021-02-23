package main

import (
	_ "medicine/common/main/configinfo"
	_ "medicine/common/main/log"
	_ "medicine/application/main/resources/db/initialize"
	_ "medicine/common/main/datasource"

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
	if err := r.Run(); err != nil {
		log.Info("startup service failed, err:%v", err)
	}
}