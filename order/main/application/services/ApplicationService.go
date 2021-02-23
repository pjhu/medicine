package orderapplicationservice

import (
	cqrs "medicine/common/main/datasource"
	log "github.com/sirupsen/logrus"
	ordercommand "medicine/order/main/application/command"
	orderresponse "medicine/order/main/application/response"
	ordermodel "medicine/order/main/domain/models"
)

// PlaceOrderHandler for create order
func PlaceOrderHandler(placeOrderCommand ordercommand.PlaceOrderCommand) (response orderresponse.PlaceOrderResponse){

	log.Info("applilcation service info:%#v\n", placeOrderCommand)
	neworder := ordermodel.PlaceOrder(6, placeOrderCommand.Quantity, "", "", "")
	
	session := cqrs.Engine.NewSession()
	defer session.Close()

	// add Begin() before any action
	err := session.Begin()
	if err != nil {
    log.Error(err)
		return
	}

	_, err = cqrs.Engine.InsertOne(&neworder)
	if err != nil {
		log.Error(err)
    session.Rollback()
    return
	}

	placeOrderResponse := orderresponse.NewPlaceOrderResponse(6)
	return placeOrderResponse
}
