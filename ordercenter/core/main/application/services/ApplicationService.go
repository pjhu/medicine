package orderapplicationservice

import (
	log "github.com/sirupsen/logrus"

	"ordercenter/common/main/errors"
	cqrs "ordercenter/common/main/datasource"
	IdGenerator "ordercenter/common/main/idgenerator"
	ordercommand "ordercenter/core/main/application/command"
	orderresponse "ordercenter/core/main/application/response"
	ordermodel "ordercenter/core/main/domain/models"
)

// PlaceOrderHandler for create order
func PlaceOrderHandler(placeOrderCommand ordercommand.PlaceOrderCommand) (result orderresponse.PlaceOrderResponse, e *errors.ErrorWithCode){

	log.Info("applilcation service info: ", placeOrderCommand)
	
	neworder := ordermodel.PlaceOrder(IdGenerator.NewId(), placeOrderCommand.Quantity, "", "", "")
	
	session := cqrs.Engine.NewSession()
	defer session.Close()

	// add Begin() before any action
	err := session.Begin()
	if err != nil {
    log.Error(err)
		return result, errors.NewErrorWithCode(errors.SystemInternalError, err.Error())
	}

	_, err = cqrs.Engine.InsertOne(&neworder)
	if err != nil {
		log.Error(err)
    session.Rollback()
    return result, errors.NewErrorWithCode(errors.SystemInternalError, "insert order error")
	}
	result = orderresponse.PlaceOrderResponse{ID: neworder.Id}
	return result, nil
}

// GetOrderDetail for get order detail
func GetOrderDetail(id int64) (rest ordermodel.UserOrder, e *errors.ErrorWithCode) {

	var order ordermodel.UserOrder 
	order = ordermodel.UserOrder{Id:id}
	has, err := cqrs.Engine.Get(&order)
	if (! has) {
		log.Error(err)
    return order, errors.NewErrorWithCode(errors.SystemInternalError, "not found order")
	}
	return order, nil
}

