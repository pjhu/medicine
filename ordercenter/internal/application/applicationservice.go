package service

import (
	"github.com/go-resty/resty/v2"
	"github.com/sirupsen/logrus"
	"ordercenter/internal/application/command"
	"ordercenter/internal/application/response"
	"ordercenter/internal/domain"
	"ordercenter/internal/pkg/errors"
	"ordercenter/internal/pkg/idgenerator"
)

type IApplicationService interface {
	PlaceOrderHandler(placeOrderCommand command.PlaceOrderCommand) (result response.PlaceOrderResponse, e *errors.ErrorWithCode)
	GetOrderDetail(id int64) (rest domain.UserOrder, e *errors.ErrorWithCode)
}

type OrderApplicationService struct {
	repository domain.IRepository
	restClient *resty.Client
}

func Build(repo domain.IRepository, client *resty.Client) IApplicationService {
	return OrderApplicationService {
		repository: repo,
		restClient: client,
	}
}

type Account struct {
	orderId string
	orderAmount int64
}

// PlaceOrderHandler for create order
func (appSvc OrderApplicationService) PlaceOrderHandler(placeOrderCommand command.PlaceOrderCommand) (result response.PlaceOrderResponse, e *errors.ErrorWithCode){

	logrus.Info("application service info: ", placeOrderCommand)
	
	newOrder := domain.PlaceOrder(idgenerator.NewId(), placeOrderCommand.Quantity, "", "", "")
	_, err := appSvc.repository.InsertOne(&newOrder)
	if err != nil {
		logrus.Error(err)
		return result, errors.NewErrorWithCode(errors.SystemInternalError, "insert order error")
	}

	post, err := appSvc.restClient.R().
		SetBody(`{"userId": 1, "orderAmount": 1}`).
		SetResult(&Account{}).
		Post("http://localhost:48080/api/v1/accounts/decrease")
	logrus.Info("account response: ",post)
	if err != nil {
		logrus.Error(err)
		return result, errors.NewErrorWithCode(errors.SystemInternalError, "decrease account error")
	}

	result = response.PlaceOrderResponse{ID: newOrder.Id}
	return result, nil
}

// GetOrderDetail for get order detail
func (appSvc OrderApplicationService) GetOrderDetail(id int64) (rest domain.UserOrder, e *errors.ErrorWithCode) {

	var order domain.UserOrder
	order = domain.UserOrder{Id: id}
	has, err := appSvc.repository.Get(&order)
	if has == nil {
		logrus.Error(err)
    	return order, errors.NewErrorWithCode(errors.SystemInternalError, "not found order")
	}
	return order, nil
}

