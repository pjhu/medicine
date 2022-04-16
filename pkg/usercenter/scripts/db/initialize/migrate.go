package dbmigrate

import (
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	_ "github.com/go-sql-driver/mysql"
    "github.com/golang-migrate/migrate/v4"
    _ "github.com/golang-migrate/migrate/v4/database/mysql"
    _ "github.com/golang-migrate/migrate/v4/source/file"
)

// Init Migrate for
func Init() {
	logrus.Info("--- Starting DB Migrate ---")
	logrus.Info(viper.GetString("datasource.master.jdbcUrl"))
	m, err := migrate.New("file://db/migrations/", "mysql://" + viper.GetString("datasource.master.jdbcUrl"))
	if err != nil {
		logrus.Error(errors.Wrap(err, "fail to start new migrate engine"))
		panic("fail to start new migrate engine")
	}
	if err := m.Up(); err != nil && err.Error() != "no change" {
		logrus.Error(err)
		panic("fail to update migration")
	}
	logrus.Info("--- Completed DB Migrate  ---")
}