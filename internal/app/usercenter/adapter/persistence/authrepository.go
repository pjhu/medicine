package persistence

import (
	"github.com/pjhu/medicine/internal/app/usercenter/domain"
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

func (r *Repo) InsertOne(member *domain.Member) error {

	if err := r.DB.Create(member).Error; err != nil {
		logrus.Error(err)
		return err
	}
	return nil
}

func (r *Repo) FindBy(member *domain.Member) error {
	if err := r.DB.Take(member).Error; err != nil {
		return err
	}
	return nil
}
