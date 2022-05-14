package persistence

import (
	"github.com/pjhu/medicine/internal/app/ordercenter/domain"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type OrderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) domain.IOrderRepository {
	return &OrderRepository{
		db: db,
	}
}

var _ domain.IOrderRepository = &OrderRepository{}

func (r *OrderRepository) InsertOne(userOrder *domain.UserOrder) error {

	if err := r.db.Create(userOrder).Error; err != nil {
		logrus.Error(err)
		return err
	}
	return nil
}

func (r *OrderRepository) FindBy(userOrder *domain.UserOrder) error {
	if err := r.db.Take(&userOrder).Error; err != nil {
		return err
	}
	return nil
}
