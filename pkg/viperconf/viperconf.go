package viperconf

import (
	"os"
	"path"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func Init(projectPath string) {
	logrus.Info("initialize viper config")
	if os.Getenv("PROFILES_ACTIVE") != "" {
		viper.SetConfigName("application" + "-" + os.Getenv("PROFILES_ACTIVE"))
	} else {
		viper.SetConfigName("application")
	}

	viper.SetConfigType("yaml")
	viper.AddConfigPath(path.Join(projectPath, "configs"))
	err := viper.ReadInConfig() // 查找并读取配置文件
	if err != nil {             // 处理读取配置文件的错误
		logrus.Error(errors.Wrap(err, "fail to read application config file"))
		panic("fail to read application config file")
	}
	logrus.Info("completed initialize viper config")
}
