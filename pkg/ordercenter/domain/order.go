package domain

import "time"

// UserOrder for Order Domain Model
type UserOrder struct {
	Id               int64 `xorm:"pk"`
	OrderAmountTotal int
	PayChannel       string    `xorm:"varchar(32)"`
	OrderStatus      string    `xorm:"varchar(25) not null"`
	CreatedAt        time.Time `xorm:"created"`
	CreatedBy        string    `xorm:"not null"`
	LastModifiedAt   time.Time `xorm:"updated"`
	LastModifiedBy   string    `xorm:"not null"`
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
