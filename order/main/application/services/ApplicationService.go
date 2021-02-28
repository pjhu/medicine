package orderapplicationservice

import (
	"errors"
	cqrs "medicine/common/main/datasource"
	IdGenerator "medicine/common/main/idgenerator"
	ordercommand "medicine/order/main/application/command"
	orderresponse "medicine/order/main/application/response"
	ordermodel "medicine/order/main/domain/models"

	log "github.com/sirupsen/logrus"
)

// PlaceOrderHandler for create order
func PlaceOrderHandler(placeOrderCommand ordercommand.PlaceOrderCommand) (result orderresponse.PlaceOrderResponse, e error){

	log.Info("applilcation service info:%#v\n", placeOrderCommand)
	
	neworder := ordermodel.PlaceOrder(IdGenerator.NewId(), placeOrderCommand.Quantity, "", "", "")
	
	session := cqrs.Engine.NewSession()
	defer session.Close()
	var response orderresponse.PlaceOrderResponse

	// add Begin() before any action
	err := session.Begin()
	if err != nil {
    log.Error(err)
		return response, err
	}

	_, err = cqrs.Engine.InsertOne(&neworder)
	if err != nil {
		log.Error(err)
    session.Rollback()
    return response, err
	}
	response = orderresponse.PlaceOrderResponse{ID: neworder.Id}
	return response, err
}

// GetOrderDetail for get order detail
func GetOrderDetail(id int64) (rest ordermodel.UserOrder, e error) {

	var order ordermodel.UserOrder 
	order = ordermodel.UserOrder{Id:id}
	has, err := cqrs.Engine.Get(&order)
	if (! has) {
		var notFoundError = errors.New("not found")
		log.Error(err)
		log.Error(notFoundError)
    return order, notFoundError
	}
	return order, nil
}

