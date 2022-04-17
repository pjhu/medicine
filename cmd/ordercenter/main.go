package main

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"pjhu/medicine/internal/app/ordercenter/routers"
	"pjhu/medicine/internal/pkg/cache"
	"pjhu/medicine/internal/pkg/datasource"
	"pjhu/medicine/pkg/httpclient"
	"pjhu/medicine/pkg/logconf"
	"pjhu/medicine/pkg/viperconf"
)

func init() {
	logconf.Init()
	viperconf.Init()
	//dbmigrate.Init()
}

func main() {
	db := datasource.BuildMysql()
	rdbRepo := cache.BuildRedis()
	restClient := httpclient.BuildRestClient()
	// 初始化路由
	newRouters := routers.Build(db, rdbRepo, restClient)
	routerEngine := newRouters.Init()

	if err := routerEngine.Run(viper.GetString("gin.port")); err != nil {
		logrus.WithError(err).Error("startup service failed")
	}
}
