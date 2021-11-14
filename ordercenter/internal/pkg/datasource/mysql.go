package datasource

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"xorm.io/xorm"
	"xorm.io/xorm/log"
	"xorm.io/xorm/names"
)

func BuildMysql() *xorm.EngineGroup {
	return initMysql()
}

// Init DBConnect for db connection
func initMysql() (engine *xorm.EngineGroup){
	conn := []string{
		viper.GetString("datasource.master.jdbcUrl"), // first one is master
		viper.GetString("datasource.slave.jdbcUrl"), // slave
	}

	var err error
	engine, err = xorm.NewEngineGroup("mysql", conn)
	if err != nil {
		logrus.Error(errors.Wrap(err, "fail to create db engine group"))
		panic("fail to create db engine group")
	}
	engine.ShowSQL(true)
	engine.SetLogLevel(log.LOG_DEBUG)
	engine.SetMapper(names.GonicMapper{})
	return engine
}
