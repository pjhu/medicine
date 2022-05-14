package persistence

import (
	"github.com/pjhu/medicine/internal/app/usercenter/domain"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type AuthRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) domain.IAuthRepository {
	return &AuthRepository{
		db: db,
	}
}

func (r *AuthRepository) InsertOne(member *domain.Member) error {

	if err := r.db.Create(member).Error; err != nil {
		logrus.Error(err)
		return err
	}
	return nil
}

func (r *AuthRepository) FindBy(member *domain.Member) error {
	if err := r.db.Take(member).Error; err != nil {
		return err
	}
	return nil
}
