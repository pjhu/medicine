package cqrs

import (
	"github.com/pkg/errors"
	_ "github.com/go-sql-driver/mysql"
	
	"xorm.io/xorm"
	xormlog "xorm.io/xorm/log"
	"xorm.io/xorm/names"

	"github.com/spf13/viper"

	log "github.com/sirupsen/logrus"
)

// Engine is EngineGroup
var Engine *xorm.EngineGroup

// DBConnect for db connection
func init() {
	conns := []string{
		viper.GetString("datasource.master.jdbcUrl"), // first one is master
		viper.GetString("datasource.slave.jdbcUrl"), // slave
	}

	var err error
	Engine, err = xorm.NewEngineGroup("mysql", conns)
	if err != nil {
		log.Error(errors.Wrap(err, "fail to create db engine group"))
		panic("fail to create db engine group")
	}
	Engine.ShowSQL(true)
	Engine.SetLogLevel(xormlog.LOG_DEBUG)
	Engine.SetMapper(names.GonicMapper{})
}
