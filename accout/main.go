package main

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"account/internal/pkg/logconf"
	"account/internal/pkg/viperconf"
	"account/routers"
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