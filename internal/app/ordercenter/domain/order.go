package domain

import "time"

// UserOrder for Order Domain Model
type UserOrder struct {
	Id               int64
	OrderAmountTotal int
	PayChannel       string
	OrderStatus      string
	CreatedAt        time.Time
	CreatedBy        string
	LastModifiedAt   time.Time
	LastModifiedBy   string
}

// PlaceOrder for create order
func PlaceOrder(id int64, orderAmountTotal int, orderStatus string, createdBy string, lastModifiedBy string) UserOrder {
	var newOrder UserOrder
	newOrder.Id = id
	newOrder.OrderAmountTotal = orderAmountTotal
	newOrder.OrderStatus = orderStatus
	newOrder.CreatedBy = createdBy
	newOrder.LastModifiedBy = lastModifiedBy
	return newOrder
}
