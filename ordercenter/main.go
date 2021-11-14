package main

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"ordercenter/internal/pkg/httpclient"

	"ordercenter/internal/pkg/cache"
	"ordercenter/internal/pkg/datasource"
	"ordercenter/internal/pkg/logconf"
	"ordercenter/internal/pkg/viperconf"
	"ordercenter/routers"
)

func init() {
	logconf.Init()
	viperconf.Init()
	cache.Init()
	//dbmigrate.Init()
}

func main() {
	db := datasource.BuildMysql()
	restClient := httpclient.BuildRestClient()
	// 初始化路由
	newRouters := routers.Build(db, restClient)
	routerEngine := newRouters.Init()

	if err := routerEngine.Run(viper.GetString("gin.port")); err != nil {
		logrus.WithError(err).Error("startup service failed")
	}
}
