package main

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"pjhu/medicine/pkg/logconf"
	"pjhu/medicine/pkg/tocbff/routers"
	"pjhu/medicine/pkg/viperconf"
)

func init() {
	logconf.Init()
	viperconf.Init()
}

func main() {
	// 初始化路由
	newRouters := routers.SetupRouterEngine()
	if err := newRouters.Run(viper.GetString("gin.port")); err != nil {
		logrus.WithError(err).Error("startup service failed")
	}
}
