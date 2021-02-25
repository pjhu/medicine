package orderapplicationservice

import (
	cqrs "medicine/common/main/datasource"
	log "github.com/sirupsen/logrus"
	ordercommand "medicine/order/main/application/command"
	orderresponse "medicine/order/main/application/response"
	ordermodel "medicine/order/main/domain/models"
	IdGenerator "medicine/common/main/idgenerator"
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
func GetOrderDetail(id int64) (result ordermodel.UserOrder, e error) {
	session := cqrs.Engine.NewSession()
	defer session.Close()

	var order ordermodel.UserOrder 
	// add Begin() before any action
	err := session.Begin()
	if err != nil {
    log.Error(err)
		return order, err
	}

	order = ordermodel.UserOrder{Id:id}
	_, err = cqrs.Engine.Get(&order)
	if err != nil {
		log.Error(err)
    session.Rollback()
    return order, err
	}
	return order, nil
}

