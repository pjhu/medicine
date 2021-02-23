package cqrs

import (
	_ "github.com/lib/pq"
	"xorm.io/xorm"
	xormlog "xorm.io/xorm/log"
	"xorm.io/xorm/names"

	log "github.com/sirupsen/logrus"
)

// Engine is EngineGroup
var Engine *xorm.EngineGroup

// DBConnect for db connection
func init() {
	conns := []string{
		"postgres://postgres:123@localhost:15432/test?sslmode=disable;", // first one is master
		"postgres://postgres:123@localhost:25432/test?sslmode=disable;", // slave
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
