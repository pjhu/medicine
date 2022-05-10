package main

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"github.com/pjhu/medicine/internal/app/usercenter/routers"
	"github.com/pjhu/medicine/internal/pkg/cache"
	"github.com/pjhu/medicine/internal/pkg/datasource"
	"github.com/pjhu/medicine/pkg/logconf"
	"github.com/pjhu/medicine/pkg/viperconf"
)

func init() {
	logconf.Init()
	viperconf.Init("cmd/usercenter")
	//dbmigrate.Build()
}

func main() {
	datasource.Builder()
	cache.Builder()
	// 初始化路由
	routerEngine := routers.Init()

	if err := routerEngine.Run(viper.GetString("gin.port")); err != nil {
		logrus.WithError(err).Error("startup service failed")
	}
}
