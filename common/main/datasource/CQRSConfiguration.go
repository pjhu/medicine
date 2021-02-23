package cqrs

import (
	_ "github.com/lib/pq"
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
	Engine, err = xorm.NewEngineGroup("postgres", conns)
	if err != nil {
		log.Error(err)
		return
	}
	Engine.ShowSQL(true)
	Engine.SetLogLevel(xormlog.LOG_DEBUG)
	Engine.SetMapper(names.GonicMapper{})
}
