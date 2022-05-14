package application

import (
	"github.com/go-resty/resty/v2"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"

	"github.com/pjhu/medicine/internal/app/ordercenter/domain"
	"github.com/pjhu/medicine/internal/pkg/cache"
	"github.com/pjhu/medicine/pkg/errors"
	"github.com/pjhu/medicine/pkg/httpclient"
	"github.com/pjhu/medicine/pkg/idgenerator"
)

type IOrderApplicationService interface {
	PlaceOrderHandler(placeOrderCommand PlaceOrderCommand) (result PlaceOrderResponse, e *errors.ErrorWithCode)
	GetOrderDetail(id int64) (rest domain.UserOrder, e *errors.ErrorWithCode)
}

type OrderApplicationService struct {
	db          *gorm.DB
	repo        domain.IOrderRepository
	cache       cache.ICacheRepository
	httprequest *resty.Request
}

// 变量断言: 如果没有实现所以方法，编译报missing method
var _ IOrderApplicationService = &OrderApplicationService{}

func NewOrderApplicationService(db *gorm.DB, repo domain.IOrderRepository,
	cache cache.ICacheRepository, httprequest *resty.Request) *OrderApplicationService {
	return &OrderApplicationService{
		db:          db,
		repo:        repo,
		cache:       cache,
		httprequest: httprequest,
	}
}

// PlaceOrderHandler for create order
func (svc *OrderApplicationService) PlaceOrderHandler(placeOrderCommand PlaceOrderCommand) (result PlaceOrderResponse, e *errors.ErrorWithCode) {

	logrus.Info("application service info: ", placeOrderCommand)
	newOrder := domain.PlaceOrder(idgenerator.NewId(), placeOrderCommand.Quantity, "", "", "")

	err := svc.db.Transaction(func(tx *gorm.DB) error {
		err := svc.repo.InsertOne(&newOrder)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		logrus.Error("place order failed, err: %v ", err)
		return result, errors.NewErrorWithCode(errors.SystemInternalError, "insert order error")
	}

	type Account struct{}
	post, err := httpclient.Request().
		SetBody(`{"userId": 1, "orderAmount": 1}`).
		SetResult(&Account{}).
		Post("http://localhost:48080/api/v1/accounts/decrease")
	logrus.Info("account response: ", post)
	if err != nil {
		logrus.Error(err)
		return result, errors.NewErrorWithCode(errors.SystemInternalError, "decrease account error")
	}

	result = PlaceOrderResponse{ID: newOrder.Id}
	return result, nil
}

// GetOrderDetail for get order detail
func (svc *OrderApplicationService) GetOrderDetail(id int64) (rest domain.UserOrder, e *errors.ErrorWithCode) {

	order := domain.UserOrder{Id: id}
	err := svc.repo.FindBy(&order)
	if err != nil {
		logrus.Error(err)
		return order, errors.NewErrorWithCode(errors.SystemInternalError, "not found order")
	}
	return order, nil
}
