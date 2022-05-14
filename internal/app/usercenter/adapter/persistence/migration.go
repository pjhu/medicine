package persistence

import (
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func Automigrate(db *gorm.DB) {
	err := db.AutoMigrate(&AuthPO{})
	if err != nil {
		logrus.Error(errors.Wrap(err, "failed to start migration"))
		panic("failed to start migration")
	}
}
