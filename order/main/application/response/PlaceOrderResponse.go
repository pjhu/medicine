package orderresponse

// PlaceOrderResponse respionse body for create order
type PlaceOrderResponse struct  {
	id int64
}

func NewPlaceOrderResponse(id int64)(response PlaceOrderResponse) {
	var newResponse PlaceOrderResponse
	newResponse.id = id
	return newResponse
}