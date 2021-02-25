package dbmigrate

import (
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

// Migrate for
func init() {
	log.Info("--- Starting DB Migrate ---")
	m, err := migrate.New("file://resources/db/migrations/", viper.GetString("datasource.master.jdbcUrl"))
	if err != nil {
		log.Error(errors.Wrap(err, "Fail to create new migrate"))
		panic("fail to creat new migrate")
	}
	if err := m.Up(); err != nil {
		log.Error(err)
	}
	log.Info("--- Completed DB Migrate  ---")
}
