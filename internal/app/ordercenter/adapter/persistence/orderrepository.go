package persistence

import (
	"github.com/pjhu/medicine/internal/app/ordercenter/domain"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Repo struct {
	DB *gorm.DB
}

func Builder(db *gorm.DB) domain.IRepository {
	return &Repo{
		DB: db,
	}
}

func (r *Repo) InsertOne(userOrder *domain.UserOrder) error {

	if err := r.DB.Create(userOrder).Error; err != nil {
		logrus.Error(err)
		return err
	}
	return nil
}

func (r *Repo) FindBy(userOrder *domain.UserOrder) error {
	if err := r.DB.Take(&userOrder).Error; err != nil {
		return err
	}
	return nil
}
