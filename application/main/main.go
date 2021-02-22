package main

import (
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	cqrs "medicine/common/main/datasource"
	log "medicine/common/main/log"
	order "medicine/order/main/adapter/rest"
)

func dbMigrate() {
	log.Info("--- Running Migration started ---")
	m, err := migrate.New("file://resources/db/migrations/", "postgres://postgres:123@localhost:15432/test?sslmode=disable")
	if err != nil {
		log.Error(err)
		return
	}
	if err := m.Up(); err != nil {
		log.Error(err)
		return
	}
	log.Info("--- Running Migration completed ---")
}

func main() {
	log.Setup()
	dbMigrate()
	cqrs.DBConnect()
	// 加载多个APP的路由配置
	Include(order.Routers)
	// 初始化路由
	r := Init()
	if err := r.Run(); err != nil {
		log.Info("startup service failed, err:%v\n", err)
	}
}
