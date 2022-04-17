package datasource

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"sync"
	"xorm.io/xorm"
	"xorm.io/xorm/log"
	"xorm.io/xorm/names"
)

var (
	once sync.Once
	mysql *xorm.EngineGroup
)

func BuildMysql() *xorm.EngineGroup {
	once.Do(func() {
		initMysql()
	})
	return mysql
}

// Init DBConnect for db connection
func initMysql() {
	conn := []string{
		viper.GetString("datasource.master.jdbcUrl"), // first one is master
		viper.GetString("datasource.slave.jdbcUrl"), // slave
	}

	var err error
	mysql, err = xorm.NewEngineGroup("mysql", conn)
	if err != nil {
		logrus.Error(errors.Wrap(err, "fail to create db engine group"))
		panic("fail to create db engine group")
	}
	mysql.ShowSQL(true)
	mysql.SetLogLevel(log.LOG_DEBUG)
	mysql.SetMapper(names.GonicMapper{})
}
