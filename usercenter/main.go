package main

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"usercenter/internal/pkg/cache"
	"usercenter/internal/pkg/datasource"
	"usercenter/internal/pkg/logconf"
	"usercenter/internal/pkg/viperconf"
	"usercenter/routers"
)

func init() {
	logconf.Init()
	viperconf.Init()
	//dbmigrate.Build()
}

func main() {
	db := datasource.BuildMysql()
	rdbRepo := cache.BuildRedis()
	// 初始化路由
	newRouters := routers.Build(db, rdbRepo)
	routerEngine := newRouters.Init()

	if err := routerEngine.Run(viper.GetString("gin.port")); err != nil {
		logrus.WithError(err).Error("startup service failed")
	}
}
