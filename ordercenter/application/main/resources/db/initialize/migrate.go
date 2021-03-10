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
		log.Error(errors.Wrap(err, "fail to start new migrate engine"))
		panic("fail to start new migrate engine")
	}
	if err := m.Up(); err != nil && err.Error() != "no change" {
		log.Error(err)
		panic("fail to update migration")
	}
	log.Info("--- Completed DB Migrate  ---")
}
