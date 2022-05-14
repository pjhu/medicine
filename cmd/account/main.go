package main

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"github.com/pjhu/medicine/internal/app/account/adapter/routers"
	"github.com/pjhu/medicine/pkg/logconf"
	"github.com/pjhu/medicine/pkg/viperconf"
)

func init() {
	logconf.Init()
	viperconf.Init()
}

func main() {
	// 初始化路由
	newRouters := routers.Init()
	if err := newRouters.Run(viper.GetString("gin.port")); err != nil {
		logrus.WithError(err).Error("startup service failed")
	}
}
