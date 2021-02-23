package viperconfiguraion

import (
	"os"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func init() {
	log.Info("initialize viper config")
	if os.Getenv("profiles_active") != "" {
		viper.SetConfigName("application" + "-" + os.Getenv("profiles_active"))
	} else {
		viper.SetConfigName("application")
	}
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./resources")
	err := viper.ReadInConfig() // 查找并读取配置文件
	if err != nil {             // 处理读取配置文件的错误

	}
	log.Info("completed initialize viper config")
}
